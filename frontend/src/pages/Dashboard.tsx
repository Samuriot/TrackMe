import { TrendingDown, TrendingUp } from "lucide-react"

import { Badge } from "@/components/ui/badge"
import { Button } from "@/components/ui/button"
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card"
import { Progress } from "@/components/ui/progress"
import { Separator } from "@/components/ui/separator"
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table"
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs"

type IncomePeriod = "weekly" | "monthly" | "yearly"

function formatMoney(amount: number) {
  return new Intl.NumberFormat("en-US", {
    style: "currency",
    currency: "USD",
    maximumFractionDigits: 0,
  }).format(amount)
}

function pct(n: number) {
  return `${Math.round(n * 100)}%`
}

export default function Dashboard() {
  const income: Record<IncomePeriod, { total: number; sources: Array<{ name: string; amount: number }> }> =
    {
      weekly: {
        total: 2350,
        sources: [
          { name: "Salary", amount: 1800 },
          { name: "Freelance", amount: 420 },
          { name: "Dividends", amount: 130 },
        ],
      },
      monthly: {
        total: 9250,
        sources: [
          { name: "Salary", amount: 7200 },
          { name: "Freelance", amount: 1400 },
          { name: "Dividends", amount: 650 },
        ],
      },
      yearly: {
        total: 118500,
        sources: [
          { name: "Salary", amount: 93600 },
          { name: "Freelance", amount: 16800 },
          { name: "Dividends", amount: 8100 },
        ],
      },
    }

  const debt = {
    total: 28450,
    items: [
      { name: "Credit Card", apr: 0.2399, balance: 3650, minPayment: 125, due: "Feb 2" },
      { name: "Car Loan", apr: 0.0649, balance: 12800, minPayment: 310, due: "Feb 10" },
      { name: "Student Loan", apr: 0.041, balance: 12000, minPayment: 180, due: "Feb 18" },
    ],
  }

  const assets = {
    cash: 12500,
    investments: 48200,
    retirement: 31200,
    homeEquity: 65000,
  }

  const totalAssets =
    assets.cash + assets.investments + assets.retirement + assets.homeEquity
  const netWorth = totalAssets - debt.total

  const stockOptions = {
    company: "Acme, Inc.",
    granted: 2000,
    vested: 850,
    strike: 4.25,
    latest409a: 9.75,
    expiry: "2033-09-30",
    vestingCliffComplete: true,
  }
  const intrinsicPerShare = Math.max(0, stockOptions.latest409a - stockOptions.strike)
  const intrinsicValue = intrinsicPerShare * stockOptions.vested
  const vestPct = (stockOptions.vested / stockOptions.granted) * 100

  const debtToAssets = debt.total / totalAssets

  return (
    <div className="min-h-screen bg-background text-foreground">
      <div className="mx-auto w-full max-w-6xl px-4 py-8">
        <div className="flex flex-col gap-3 sm:flex-row sm:items-end sm:justify-between">
          <div>
            <div className="text-3xl font-semibold tracking-tight text-h1">Dashboard</div>
            <p className="text-sm text-h2">
              POC for a dashboard using shadcn
            </p>
          </div>
          <div className="flex items-center gap-2">
            <Button variant="outline" className="text-h2">Add transaction</Button>
            <Button>Sync accounts</Button>
          </div>
        </div>

        <Separator className="my-6" />

        <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-4">
          <Card>
            <CardHeader className="pb-2">
              <CardDescription>Total income (monthly)</CardDescription>
              <CardTitle className="text-2xl">{formatMoney(income.monthly.total)}</CardTitle>
            </CardHeader>
            <CardContent className="pt-0">
              <div className="flex items-center gap-2 text-sm text-muted-foreground">
                <TrendingUp className="h-4 w-4" />
                +8% vs last month
              </div>
            </CardContent>
          </Card>

          <Card>
            <CardHeader className="pb-2">
              <CardDescription>Total debt</CardDescription>
              <CardTitle className="text-2xl">{formatMoney(debt.total)}</CardTitle>
            </CardHeader>
            <CardContent className="pt-0">
              <div className="flex items-center gap-2 text-sm text-muted-foreground">
                <TrendingDown className="h-4 w-4" />
                -2.4% vs last month
              </div>
            </CardContent>
          </Card>

          <Card>
            <CardHeader className="pb-2">
              <CardDescription>Net worth</CardDescription>
              <CardTitle className="text-2xl">{formatMoney(netWorth)}</CardTitle>
            </CardHeader>
            <CardContent className="pt-0">
              <div className="text-sm text-muted-foreground">
                Assets: {formatMoney(totalAssets)}
              </div>
            </CardContent>
          </Card>

          <Card>
            <CardHeader className="pb-2">
              <CardDescription>Options (vested value)</CardDescription>
              <CardTitle className="text-2xl">{formatMoney(intrinsicValue)}</CardTitle>
            </CardHeader>
            <CardContent className="pt-0">
              <div className="flex items-center justify-between text-sm text-muted-foreground">
                <span>
                  {stockOptions.vested}/{stockOptions.granted} vested
                </span>
                <span>{Math.round(vestPct)}%</span>
              </div>
              <Progress className="mt-2" value={vestPct} />
            </CardContent>
          </Card>
        </div>

        <div className="mt-8 grid gap-6 lg:grid-cols-3">
          <Card className="lg:col-span-2">
            <CardHeader>
              <CardTitle>Income</CardTitle>
              <CardDescription>Weekly / monthly / yearly totals and breakdown.</CardDescription>
            </CardHeader>
            <CardContent>
              <Tabs defaultValue="monthly" className="w-full">
                <TabsList>
                  <TabsTrigger value="weekly">Weekly</TabsTrigger>
                  <TabsTrigger value="monthly">Monthly</TabsTrigger>
                  <TabsTrigger value="yearly">Yearly</TabsTrigger>
                </TabsList>

                {(["weekly", "monthly", "yearly"] as const).map((period) => {
                  const total = income[period].total
                  return (
                    <TabsContent key={period} value={period}>
                      <div className="flex flex-col gap-4">
                        <div className="flex items-end justify-between">
                          <div>
                            <div className="text-sm text-muted-foreground">Total</div>
                            <div className="text-2xl font-semibold">
                              {formatMoney(total)}
                            </div>
                          </div>
                          <Badge variant="secondary">dummy</Badge>
                        </div>

                        <div className="space-y-3">
                          {income[period].sources.map((s) => {
                            const share = s.amount / total
                            return (
                              <div key={s.name} className="space-y-1">
                                <div className="flex items-center justify-between text-sm">
                                  <span className="font-medium">{s.name}</span>
                                  <span className="text-muted-foreground">
                                    {formatMoney(s.amount)} · {pct(share)}
                                  </span>
                                </div>
                                <Progress value={share * 100} />
                              </div>
                            )
                          })}
                        </div>
                      </div>
                    </TabsContent>
                  )
                })}
              </Tabs>
            </CardContent>
          </Card>

          <Card>
            <CardHeader>
              <CardTitle>Balance health</CardTitle>
              <CardDescription>Quick ratios to watch.</CardDescription>
            </CardHeader>
            <CardContent className="space-y-5">
              <div className="space-y-2">
                <div className="flex items-center justify-between text-sm">
                  <span className="text-muted-foreground">Debt-to-assets</span>
                  <span className="font-medium">{pct(debtToAssets)}</span>
                </div>
                <Progress value={Math.min(100, debtToAssets * 100)} />
                <p className="text-xs text-muted-foreground">
                  Lower is generally better. This is a rough heuristic.
                </p>
              </div>

              <div className="rounded-lg border bg-muted/30 p-4">
                <div className="text-sm font-medium">Assets snapshot</div>
                <div className="mt-3 space-y-2 text-sm">
                  <div className="flex items-center justify-between">
                    <span className="text-muted-foreground">Cash</span>
                    <span>{formatMoney(assets.cash)}</span>
                  </div>
                  <div className="flex items-center justify-between">
                    <span className="text-muted-foreground">Investments</span>
                    <span>{formatMoney(assets.investments)}</span>
                  </div>
                  <div className="flex items-center justify-between">
                    <span className="text-muted-foreground">Retirement</span>
                    <span>{formatMoney(assets.retirement)}</span>
                  </div>
                  <div className="flex items-center justify-between">
                    <span className="text-muted-foreground">Home equity</span>
                    <span>{formatMoney(assets.homeEquity)}</span>
                  </div>
                </div>
              </div>
            </CardContent>
          </Card>
        </div>

        <div className="mt-8 grid gap-6 lg:grid-cols-2">
          <Card>
            <CardHeader>
              <CardTitle>Debt</CardTitle>
              <CardDescription>Balances, APR, and upcoming minimums.</CardDescription>
            </CardHeader>
            <CardContent>
              <Table>
                <TableHeader>
                  <TableRow>
                    <TableHead>Account</TableHead>
                    <TableHead>APR</TableHead>
                    <TableHead className="text-right">Balance</TableHead>
                    <TableHead className="text-right">Min</TableHead>
                    <TableHead className="text-right">Due</TableHead>
                  </TableRow>
                </TableHeader>
                <TableBody>
                  {debt.items.map((d) => (
                    <TableRow key={d.name}>
                      <TableCell className="font-medium">{d.name}</TableCell>
                      <TableCell>{pct(d.apr)}</TableCell>
                      <TableCell className="text-right">{formatMoney(d.balance)}</TableCell>
                      <TableCell className="text-right">{formatMoney(d.minPayment)}</TableCell>
                      <TableCell className="text-right">
                        <Badge variant="outline">{d.due}</Badge>
                      </TableCell>
                    </TableRow>
                  ))}
                </TableBody>
              </Table>
              <div className="mt-4 text-sm text-muted-foreground">
                Total debt: <span className="font-medium text-foreground">{formatMoney(debt.total)}</span>
              </div>
            </CardContent>
          </Card>

          <Card>
            <CardHeader>
              <CardTitle>Stock options</CardTitle>
              <CardDescription>Vesting and intrinsic value (409A - strike).</CardDescription>
            </CardHeader>
            <CardContent className="space-y-5">
              <div className="flex flex-wrap items-center gap-2">
                <Badge>{stockOptions.company}</Badge>
                <Badge variant="secondary">
                  Strike: {formatMoney(stockOptions.strike)}
                </Badge>
                <Badge variant="secondary">
                  409A: {formatMoney(stockOptions.latest409a)}
                </Badge>
                {stockOptions.vestingCliffComplete ? (
                  <Badge variant="outline">Cliff complete</Badge>
                ) : (
                  <Badge variant="outline">In cliff</Badge>
                )}
              </div>

              <div className="space-y-2">
                <div className="flex items-center justify-between text-sm">
                  <span className="text-muted-foreground">Vested</span>
                  <span className="font-medium">
                    {stockOptions.vested}/{stockOptions.granted} ({Math.round(vestPct)}%)
                  </span>
                </div>
                <Progress value={vestPct} />
              </div>

              <div className="grid gap-3 rounded-lg border bg-muted/30 p-4 sm:grid-cols-2">
                <div>
                  <div className="text-xs text-muted-foreground">Intrinsic / share</div>
                  <div className="text-lg font-semibold">{formatMoney(intrinsicPerShare)}</div>
                </div>
                <div>
                  <div className="text-xs text-muted-foreground">Vested value</div>
                  <div className="text-lg font-semibold">{formatMoney(intrinsicValue)}</div>
                </div>
                <div className="sm:col-span-2">
                  <div className="text-xs text-muted-foreground">Expiry</div>
                  <div className="text-sm font-medium">{stockOptions.expiry}</div>
                </div>
              </div>
            </CardContent>
          </Card>
        </div>
      </div>
    </div>
  )
}
