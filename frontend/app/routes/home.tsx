import type { Route } from "./+types/home";
import { Dashboard } from "../components/Dashboard";

export function meta({ }: Route.MetaArgs) {
  return [
    { title: "Dogs API" },
    { name: "description", content: "Welcome to Dogs API!" },
  ];
}

export default function Home() {
  return <Dashboard />;
}
