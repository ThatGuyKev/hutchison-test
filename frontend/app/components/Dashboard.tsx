import { Typography } from "@mui/material";
import { useEffect, useState } from "react";
import SearchBar from "./SearchBar";
import TableDogs from "./TableDogs";
import { listDogs } from "api";



export type DogIndex = {
  id: number;
  breed: string;
  variants: string[];
  _searchText: string;
};

const search = (index: DogIndex[], query: string) => {
  return index.filter((item: DogIndex) =>
    item._searchText.includes(query.toLowerCase())
  );
};

export function Dashboard() {
  const [dogsIndex, setDogsIndex] = useState([] as DogIndex[]);
  const [filteredList, setFilteredList] = useState([] as DogIndex[]);

  const fetchDogs = async () => {
    console.log("Fetching dogs")
    let dogs = await listDogs()

    let index = dogs.map((dog => {
      const searchText = dog.Variants + dog.Breed
      const variants: string[] = JSON.parse(dog.Variants)
      return {
        id: dog.ID,
        breed: dog.Breed,
        variants: variants,
        _searchText: searchText
      }
    }))
    setDogsIndex(index)
    setFilteredList(index)


  }
  useEffect(() => {
    fetchDogs()
  }, []);

  const fullTextSearch = (q: string) => {
    let newList = search(dogsIndex, q)
    setFilteredList(newList);
  };

  return (
    <main>
      <div>
        <header className="flex flex-col items-center gap-9">
          <div className="w-[500px] max-w-[100vw] p-4">
            <Typography
              autoCapitalize="words"
              gutterBottom
              variant="h1"
              component="div"
            >
              Dogs API
            </Typography>
          </div>
        </header>


        <SearchBar onChange={fullTextSearch} />
          <TableDogs data={filteredList} refetch={fetchDogs}/>
      </div>
    </main>
  );
}
