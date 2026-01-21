import { Link } from "react-router-dom"

import { Button } from "@/components/ui/button"

function Home() {
  return (
    <div className="min-h-screen bg-background text-foreground">
      <div className="mx-auto flex w-full max-w-4xl flex-col gap-6 px-4 py-12">
        <div className="space-y-2">
          <h1 className="text-4xl font-semibold tracking-tight">TrackMe</h1>
          <p className="text-muted-foreground">
            A simple finance tracker proof-of-concept.
          </p>
        </div>

        <div className="flex items-center gap-3">
          <Button asChild>
            <Link to="/dashboard">Go to dashboard</Link>
          </Button>
          <Button asChild variant="outline">
            <a href="https://ui.shadcn.com" target="_blank" rel="noreferrer">
              shadcn/ui
            </a>
          </Button>
        </div>

        <div className="rounded-xl border bg-card p-6 text-card-foreground">
          <div className="text-sm text-muted-foreground">
            Tip: Your global background comes from <code>bg-background</code>,
            which is driven by <code>--background</code> in <code>index.css</code>.
          </div>
        </div>
      </div>
    </div>
  )
}

export default Home
