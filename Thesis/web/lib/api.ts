const BASE = process.env.NEXT_PUBLIC_API_BASE_URL?.replace(/\/$/, '') || 'http://localhost:8080'


export type IngestResp = { collection: string; chunks: number; title: string }
export type SummarizeResp = { heading: string; summary: string; evidence: any[] }


export async function ingest(file: File): Promise<IngestResp> {
const fd = new FormData()
fd.append('file', file)
const res = await fetch(`${BASE}/ingest`, { method: 'POST', body: fd })
if (!res.ok) throw new Error(`ingest failed: ${res.status}`)
return res.json()
}


export async function summarize(collection: string, heading: string, topK = 10): Promise<SummarizeResp> {
const res = await fetch(`${BASE}/summarize`, {
method: 'POST',
headers: { 'Content-Type': 'application/json' },
body: JSON.stringify({ collection, heading, top_k: topK }),
})
if (!res.ok) throw new Error(`summarize failed: ${res.status}`)
return res.json()
}