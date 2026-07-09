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

    async function handleSubmit(
        e: React.FormEvent<HTMLFormElement>
    ) {

        e.preventDefault();

        setError("");
        setLoading(true);

        try {

            const response = await createShortURL({
                url,
                alias,
            });

            onSuccess(response);

            setURL("");
            setAlias("");

        } catch (err) {

            if (err instanceof Error) {
                setError(err.message);
            } else {
                setError("Something went wrong.");
            }

        } finally {

            setLoading(false);

        }

    }

    return (

        <form
            onSubmit={handleSubmit}
            className="space-y-5"
        >

            <div>

                <label className="mb-2 block text-sm font-medium text-slate-300">
                    Original URL
                </label>

                <input
                    type="url"
                    placeholder="https://example.com"
                    value={url}
                    onChange={(e) => setURL(e.target.value)}
                    className="
                        w-full
                        rounded-xl
                        border
                        border-slate-700
                        bg-slate-900/80
                        px-4
                        py-3
                        text-white
                        outline-none
                        transition
                        placeholder:text-slate-500
                        focus:border-blue-500
                        focus:ring-2
                        focus:ring-blue-500/20
                    "
                    required
                />

            </div>

            <div>

                <label className="mb-2 block text-sm font-medium text-slate-300">
                    Custom Alias
                </label>

                <input
                    type="text"
                    placeholder="Optional"
                    value={alias}
                    onChange={(e) => setAlias(e.target.value)}
                    className="
                        w-full
                        rounded-xl
                        border
                        border-slate-700
                        bg-slate-900/80
                        px-4
                        py-3
                        text-white
                        outline-none
                        transition
                        placeholder:text-slate-500
                        focus:border-blue-500
                        focus:ring-2
                        focus:ring-blue-500/20
                    "
                />

            </div>

            {error && (

                <div className="rounded-xl border border-red-500/30 bg-red-500/10 px-4 py-3 text-sm text-red-300">

                    {error}

                </div>

            )}

            <button
                type="submit"
                disabled={loading}
                className="
                    w-full
                    rounded-xl
                    bg-blue-600
                    py-3
                    font-semibold
                    text-white
                    transition-all
                    duration-200
                    hover:bg-blue-700
                    active:scale-[0.98]
                    disabled:cursor-not-allowed
                    disabled:opacity-60
                "
            >

                {loading
                    ? "Shortening..."
                    : "Shorten URL"}

            </button>

        </form>

    );

}
