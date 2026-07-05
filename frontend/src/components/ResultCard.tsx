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

        <div className="mt-8 rounded-2xl border border-slate-700 bg-slate-900 p-6">

            <h2 className="mb-4 text-lg font-semibold text-white">
                ✅ Short URL Created
            </h2>

            <div className="space-y-2">

                <p className="text-sm text-slate-400">
                    Original URL
                </p>

                <p className="break-all text-white">
                    {result.original_url}
                </p>

            </div>

            <div className="mt-6 flex gap-3">

                <input
                    readOnly
                    value={result.short_url}
                    className="flex-1 rounded-xl border border-slate-700 bg-slate-800 px-4 py-3 text-white"
                />

                <button
                    onClick={handleCopy}
                    className="rounded-xl bg-green-600 px-5 py-3 font-semibold text-white transition hover:bg-green-700"
                >
                    {copied ? "Copied!" : "Copy"}
                </button>

            </div>

        </div>

    );

}
