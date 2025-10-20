var baseUrl = process.env.BACKEND_BASE_URL || "https://hutchison-test.onrender.com/api";
interface Dog {
  ID: number;
  Breed: string;
  Variants: string;
  CreatedAt: string;
  UpdatedAt: string;
  DeletedAt: string;
}


export async function createDog(breed: string, variants: string): Promise<Dog> {
  const body = {
    Breed: breed,
    Variants: variants,
  };
  const res = await fetch(`${baseUrl}/dogs`, {
    method: "Post",
    body: JSON.stringify(body),
  });

  console.log(res);
  const jsonRes = await res.json();

  return jsonRes.Data.Dog;
}
export async function listDogs(): Promise<Dog[]> {
  console.log("listing dogs")
  let res = await fetch(baseUrl + "/dogs");

  const jsonRes = await res.json();
  const dogs = jsonRes.Data.Dogs;
  return dogs;
}

export async function editDog(id: number,breed: string, variants: string): Promise<Dog> {
   const body = {
    Breed: breed,
    Variants:variants,
  };
  const res = await fetch(`${baseUrl}/dogs/${id}`, {
    method: "Put",
    body: JSON.stringify(body),
  });

  const jsonRes = await res.json();
  console.log(jsonRes);

  return jsonRes.Data.Dog;
}

export async function deleteDog(id: number): Promise<boolean> {
  const res = await fetch(`${baseUrl}/dogs/${id}`, { method: "DELETE" });
  console.log(res);
  const jsonRes = await res.json()
  return jsonRes.ok;
}