export default function Footer() {
    return (

        <footer className="mt-12 flex flex-wrap items-center justify-center gap-4">

            <span className="rounded-full bg-slate-800 px-4 py-2 text-sm text-slate-300">
                Go
            </span>

            <span className="rounded-full bg-slate-800 px-4 py-2 text-sm text-slate-300">
                PostgreSQL
            </span>

            <span className="rounded-full bg-slate-800 px-4 py-2 text-sm text-slate-300">
                Docker
            </span>

            <a
                href="https://github.com/shantam-sharma/url-shortener"
                target="_blank"
                rel="noopener noreferrer"
                className="rounded-full bg-slate-700 px-4 py-2 text-sm font-medium text-white transition-all hover:bg-slate-600"
            >
                GitHub →
            </a>

        </footer>

    );
}
