import * as React from "react";
import {
  Box, Table, TableBody, TableCell, TableContainer,
  TableHead,
  TablePagination,
  TableRow,
  Toolbar,
  Typography,
  Paper,
  Tooltip,
  Chip,
  Dialog,
  DialogActions,
  DialogTitle,
  DialogContentText,
  Button,
  DialogContent,
  Snackbar,
  TextField,
  FormControl,
} from "@mui/material";

import { EditNote as EditNoteIcon, DeleteForever as DeleteForeverIcon, AddBox as AddIcon } from "@mui/icons-material";
import type { DogIndex } from "./Dashboard";
import { createDog, deleteDog, editDog } from "api";
import { useState } from "react";


interface HeadCell {
  disablePadding: boolean;
  id: string;
  label: string;
  align: "left" | "right";
}

const headCells: readonly HeadCell[] = [
  {
    id: "breed",
    align: "left",
    disablePadding: true,
    label: "Breed",
  },
  {
    id: "variants",
    align: "left",
    disablePadding: false,
    label: "Variants",
  },
  {
    id: "actions",
    align: "right",
    disablePadding: false,
    label: "Actions",
  }
];



export default function TableDogs({ data, refetch }: { data: DogIndex[], refetch: () => void }) {

  const [page, setPage] = useState(0);
  const [deleting, setDeleting] = useState({ isDeleting: false, id: 0, breed: "" });
  const [rowsPerPage, setRowsPerPage] = useState(5);

  const [isCreating, setIsCreating] = useState(false);


  const [isEditing, setIsEditing] = useState(false)
  const [snackbarState, setSnackbarState] = useState({ isOpen: false, message: "" })

  const handleDelete = async (id: number, breed: string) => {
    console.log("Deleting dog with id:", id);
    let res = await deleteDog(id)
    if (res) {
      setSnackbarState({ isOpen: true, message: `Dog ${breed} deleted successfully!` })
    } else {
      setSnackbarState({ isOpen: true, message: `Issue deleting ${breed}!` })

    }
    refetch()
    setDeleting({ isDeleting: false, id: 0, breed: "" })

  }

  const handleEditDog = async (id: number, breed: string, variants: string) => {
    if (id == 0) {
      setSnackbarState({ isOpen: true, message: "Issue editing dog" })
      setSelectedDog(null)
      setIsEditing(false)
      return
    }

    let res = await editDog(id, breed, variants)
    if (res) {
      setSnackbarState({ isOpen: true, message: "Dog edited successfully" })
    } else {
      setSnackbarState({ isOpen: true, message: "Issue editing dog" })

    }
    refetch()
    setIsEditing(false)


  }

  const handleChangePage = (event: unknown, newPage: number) => {
    setPage(newPage);
  };

  const handleChangeRowsPerPage = (
    event: React.ChangeEvent<HTMLInputElement>
  ) => {
    setRowsPerPage(parseInt(event.target.value, 10));
    setPage(0);
  };

  const emptyRows =
    page > 0 ? Math.max(0, (1 + page) * rowsPerPage - data.length) : 0;

  const visibleRows = [...data].slice(page * rowsPerPage, page * rowsPerPage + rowsPerPage)

  const handleCreateDog = async (newBreed: string, newVariants: string) => {

    console.log(`Creating new Dog ${newBreed} | ${newVariants}`)

    let res = await createDog(newBreed, newVariants)

    if (res) {
      setSnackbarState({ isOpen: true, message: "New Dog created successfully" })
    }
    refetch()
  }


  const [selectedDog, setSelectedDog] = useState<DogIndex | null>(null)

  const CreateOrEditDialog = () => {

    const [newBreed, setNewBreed] = useState(selectedDog?.breed || "");
    const [newVariants, setNewVariants] = useState(selectedDog?.variants?.join(", ") || "");
    return <Dialog
      open={isCreating || isEditing}
      onClose={() => {
        setSelectedDog(null)
        setIsCreating(false)
        setIsEditing(false)
      }}
      aria-labelledby="alert-dialog-title"
      aria-describedby="alert-dialog-description"
    >
      <DialogTitle id="alert-dialog-title">
        {isEditing ? "Edit Dog" : "Create New Dog"}
      </DialogTitle>
      <DialogContent>
        <FormControl>
          <TextField sx={{ margin: 1 }} label="Breed" variant="outlined" value={newBreed} onChange={(e) => setNewBreed(e.target.value)} />
          <TextField sx={{ margin: 1 }} label="Variants (comma separated)" variant="outlined" value={newVariants} onChange={(e) => setNewVariants(e.target.value)} />
        </FormControl>
      </DialogContent>
      <DialogActions>
        <Button
          onClick={() => { setIsCreating(false); setIsEditing(false); setSelectedDog(null) }}>Cancel</Button>
        <Button
          onClick={() => {
            if (isCreating) {
              handleCreateDog(newBreed, newVariants)
            }
            else {
              handleEditDog(selectedDog?.id || 0, newBreed, newVariants)
            }
          }} autoFocus>
          {isCreating ? "Create" : "Save"}
        </Button>
      </DialogActions>
    </Dialog>
  }

  return (
    <Box component="section" sx={{ p: 4 }}>
      <Paper sx={{ width: "100%", mb: 2, p: 4 }}>
        <Toolbar sx={[{
          pl: { sm: 2 },
          pr: { xs: 1, sm: 1 },
        },]}
        >
          <Typography
            sx={{ flex: "1 1 100%" }}
            variant="h6"
            id="tableTitle"
            component="div"
          >
            Dogs
          </Typography>

          <Tooltip title="Filter list">
            <Button onClick={() => setIsCreating(true)}>
              <AddIcon />
            </Button>
          </Tooltip>

        </Toolbar>

        <TableContainer>
          <Table
            sx={{ minWidth: 750 }}
            aria-labelledby="tableTitle"
            size="small"
          >
            <TableHead>
              <TableRow>
                {headCells.map((headCell) => (
                  <TableCell
                    key={headCell.id}
                    align={headCell.align}
                    padding={headCell.disablePadding ? "none" : "normal"}
                  >{headCell.label}</TableCell>
                ))}
              </TableRow>
            </TableHead>
            <TableBody>
              {visibleRows.map((row, index) => {
                const labelId = `enhanced-table-checkbox-${index}`;

                return (
                  <TableRow
                    role="checkbox"
                    tabIndex={-1}
                    key={row.id}
                  >
                    <TableCell
                      component="th"
                      id={labelId}
                      scope="row"
                      padding="none"
                      style={{ textTransform: "capitalize" }}
                    >
                      {row.breed}
                    </TableCell>
                    <TableCell align="left">
                      {row.variants?.map(variant =>
                        <Chip sx={{ margin: 0.5 }} key={variant} label={variant} size="small" />
                      )}

                    </TableCell>
                    <TableCell align="right">
                      <Button>
                        <EditNoteIcon color="action" onClick={() => {
                          setIsEditing(true)
                          setSelectedDog(row)
                        }} />
                      </Button>
                      <Button><DeleteForeverIcon color="error" onClick={() => setDeleting({ isDeleting: true, id: row.id, breed: row.breed })} /></Button>

                    </TableCell>
                  </TableRow>
                );
              })}
              {emptyRows > 0 && (
                <TableRow
                  style={{
                    height: 33 * emptyRows,
                  }}
                >
                  <TableCell colSpan={6} />
                </TableRow>
              )}
            </TableBody>
          </Table>
        </TableContainer>
        <TablePagination
          rowsPerPageOptions={[5, 10, 25]}
          component="div"
          count={data.length}
          rowsPerPage={rowsPerPage}
          page={page}
          onPageChange={handleChangePage}
          onRowsPerPageChange={handleChangeRowsPerPage}
        />
      </Paper>
      <Dialog
        open={deleting.isDeleting}
        onClose={() => console.log("deleting")}
        aria-labelledby="alert-dialog-title"
        aria-describedby="alert-dialog-description"
      >
        <DialogTitle id="alert-dialog-title">
          {"Deleting.."}
        </DialogTitle>
        <DialogContent>
          <DialogContentText id="alert-dialog-description">
            Are you sure you want to delete dog?<br />
            {deleting.breed}
          </DialogContentText>
        </DialogContent>
        <DialogActions>
          <Button
            onClick={() => setDeleting({ isDeleting: false, id: 0, breed: "" })}>Cancel</Button>
          <Button
            onClick={() => handleDelete(deleting.id, deleting.breed)} autoFocus>
            Delete
          </Button>
        </DialogActions>
      </Dialog>

      <CreateOrEditDialog />
      <Snackbar
        open={snackbarState.isOpen}
        autoHideDuration={3000}
        onClose={() => setSnackbarState({ isOpen: false, message: "" })}
        message={snackbarState.message} />



      <Snackbar
        open={snackbarState.isOpen}
        autoHideDuration={3000}
        onClose={() => setSnackbarState({ isOpen: false, message: "" })}
        message={snackbarState.message} />

    </Box >
  );
}
