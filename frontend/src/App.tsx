import { useState } from "react";

import Header from "./components/Header";
import URLForm from "./components/URLForm";
import ResultCard from "./components/ResultCard";
import Footer from "./components/Footer";

import type { CreateURLResponse } from "./types/api";

export default function App() {

    const [result, setResult] = useState<CreateURLResponse | null>(null);

    return (

        <main className="relative min-h-screen overflow-hidden bg-slate-950 px-6 py-12">

            {/* Background Glow */}
            <div className="pointer-events-none absolute -left-40 -top-40 h-96 w-96 rounded-full bg-blue-500/10 blur-3xl" />
            <div className="pointer-events-none absolute -bottom-40 -right-40 h-96 w-96 rounded-full bg-violet-500/10 blur-3xl" />

            <div className="relative mx-auto max-w-2xl">

                <Header />

                <div
                    className="
                        rounded-3xl
                        border
                        border-slate-700/60
                        bg-slate-900/80
                        p-8
                        shadow-2xl
                        backdrop-blur-sm
                    "
                >

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
