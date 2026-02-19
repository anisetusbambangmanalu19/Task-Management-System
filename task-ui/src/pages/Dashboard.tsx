import { useEffect, useState } from "react";
import api from "../services/api";

export default function Dashboard() {
  const [projects, setProjects] = useState<any[]>([]);
  const [name, setName] = useState("");

  const fetchProjects = async () => {
    const res = await api.get("/projects");
    console.log("PROJECTS DATA:", res.data);
    setProjects(res.data);
};


  const createProject = async () => {
    if (!name) return;
    await api.post("/projects", { name, description: "" });
    setName("");
    fetchProjects();
  };

  useEffect(() => {
    fetchProjects();
  }, []);

  return (
    <div className="min-h-screen bg-gray-100 p-10">
      <div className="max-w-3xl mx-auto bg-white shadow-lg rounded-xl p-6">
        <h1 className="text-2xl font-bold mb-6">My Projects</h1>

        <div className="flex gap-2 mb-6">
          <input
            className="flex-1 border rounded-lg px-4 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500"
            placeholder="New Project"
            value={name}
            onChange={(e) => setName(e.target.value)}
          />
          <button
            onClick={createProject}
            className="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition"
          >
            Add
          </button>
        </div>

        <div className="space-y-3">
          {projects.map((p) => (
            <div
            key={p.id}
            className="border rounded-lg p-4 bg-white shadow-sm"
            >
            <h2 className="font-semibold text-lg text-gray-800">
            {p.Name}
            </h2>
            </div>

          ))}
        </div>
      </div>
    </div>
  );
}
