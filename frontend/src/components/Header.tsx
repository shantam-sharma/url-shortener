export default function Header() {
    return (

        <header className="mb-12 text-center">

            <div className="mb-5 inline-flex items-center rounded-full border border-slate-700 bg-slate-900/70 px-4 py-2 text-sm text-slate-300 backdrop-blur-sm">

                🚀 Built with Go • PostgreSQL • React

            </div>

            <h1 className="text-5xl font-bold tracking-tight text-white">

                Go URL Shortener

            </h1>

            <p className="mx-auto mt-5 max-w-xl text-lg leading-7 text-slate-400">

                A fast and reliable URL shortening service built with
                Go, PostgreSQL, Docker and React.

            </p>

        </header>



    );
}
