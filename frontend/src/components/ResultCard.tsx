import { useState } from "react";

import type { CreateURLResponse } from "../types/api";

interface ResultCardProps {
    result: CreateURLResponse;
}

export default function ResultCard({ result }: ResultCardProps) {

    const [copied, setCopied] = useState(false);

    async function handleCopy() {

        await navigator.clipboard.writeText(result.short_url);

        setCopied(true);

        setTimeout(() => {
            setCopied(false);
        }, 2000);

    }

    return (

        <div className="mt-8 rounded-2xl border border-slate-700/70 bg-slate-900/80 p-6 shadow-lg backdrop-blur-sm transition-all">

            <h2 className="mb-5 text-xl font-semibold text-green-400">
                ✓ Short URL Created
            </h2>

            <div className="space-y-2">

                <p className="text-sm font-medium text-slate-400">
                    Short URL
                </p>

                <div className="flex gap-3">

                    <input
                        readOnly
                        value={result.short_url}
                        className="flex-1 rounded-xl border border-slate-700 bg-slate-800 px-4 py-3 text-white outline-none"
                    />

                    <button
                        onClick={handleCopy}
                        className="rounded-xl bg-green-600 px-5 py-3 font-semibold text-white transition-all duration-200 hover:bg-green-700 active:scale-95"
                    >
                        {copied ? "✓ Copied" : "Copy"}
                    </button>

                </div>

            </div>

            <div className="mt-6 border-t border-slate-700 pt-6">

                <p className="mb-2 text-sm font-medium text-slate-400">
                    Original URL
                </p>

                <p className="break-all rounded-xl bg-slate-800 px-4 py-3 text-slate-200">
                    {result.original_url}
                </p>

            </div>

        </div>

    );

}
