import { useState } from "react";

import { createShortURL } from "../services/api";
import type { CreateURLResponse } from "../types/api";

interface URLFormProps {
    onSuccess: (data: CreateURLResponse) => void;
}

export default function URLForm({ onSuccess }: URLFormProps) {

    const [url, setURL] = useState("");
    const [alias, setAlias] = useState("");

    const [loading, setLoading] = useState(false);
    const [error, setError] = useState("");

    const handleSubmit = async (
        e: React.FormEvent<HTMLFormElement>
    ) => {

        e.preventDefault();
        setError("");
        setLoading(true);

        try {

            const response = await createShortURL({
                url,
                alias,
            });

            setError("");

            onSuccess(response);

            setURL("");
            setAlias("");

        } catch (error) {

            if (error instanceof Error) {
                setError(error.message);
            } else {
                setError("Something went wrong.");
            }

        } finally {

            setLoading(false);

        }

    };

    return (

        <form
            onSubmit={handleSubmit}
            className="space-y-5"
        >

            <input
                type="url"
                placeholder="https://example.com"
                value={url}
                onChange={(e) => setURL(e.target.value)}
                className="w-full rounded-xl border border-slate-700 bg-slate-900 px-4 py-3 text-white outline-none transition focus:border-blue-500"
                required
            />

            <input
                type="text"
                placeholder="Custom alias (optional)"
                value={alias}
                onChange={(e) => setAlias(e.target.value)}
                className="w-full rounded-xl border border-slate-700 bg-slate-900 px-4 py-3 text-white outline-none transition focus:border-blue-500"
            />

            {error && (
                <div className="rounded-xl border border-red-500/30 bg-red-500/10 px-4 py-3 text-sm text-red-300">
                    {error}
                </div>
            )}

            <button
                type="submit"
                disabled={loading}
                className="w-full rounded-xl bg-blue-600 py-3 font-semibold text-white transition hover:bg-blue-700 disabled:cursor-not-allowed disabled:opacity-50"
            >

                {loading
                    ? "Shortening..."
                    : "Shorten URL"}

            </button>

        </form>

    );

}
