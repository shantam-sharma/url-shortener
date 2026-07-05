import { useState } from "react";

import Header from "./components/Header";
import URLForm from "./components/URLForm";
import ResultCard from "./components/ResultCard";
import Footer from "./components/Footer";

import type { CreateURLResponse } from "./types/api";

export default function App() {

    const [result, setResult] = useState<CreateURLResponse | null>(null);

    return (

        <main className="min-h-screen bg-slate-950 px-6 py-12">

            <div className="mx-auto max-w-2xl">

                <Header />

                <div className="rounded-3xl border border-slate-800 bg-slate-900 p-8 shadow-2xl">

                    <URLForm
                        onSuccess={setResult}
                    />

                    {result && (
                        <ResultCard
                            result={result}
                        />
                    )}

                </div>

                <Footer />

            </div>

        </main>

    );

}
