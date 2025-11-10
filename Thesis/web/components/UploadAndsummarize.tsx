'use client'
<div className="flex flex-col sm:flex-row gap-3 items-start sm:items-center">
<input type="file" accept="application/pdf" onChange={e=>onSelect(e.target.files?.[0] ?? null)} />
<button disabled={!canIngest} onClick={doIngest}
className="inline-flex items-center gap-2 rounded-lg bg-black px-4 py-2 text-white disabled:opacity-50">
{loading==='ingest' && <Loader2 className="h-4 w-4 animate-spin"/>}
Ingest
</button>
</div>
{title && (
<p className="mt-3 text-sm text-gray-600">Title: <span className="font-medium text-gray-800">{title}</span> / Chunks: {chunks}</p>
)}
</section>


<section className="rounded-2xl border bg-white p-5 shadow-sm">
<div className="flex items-center gap-2 mb-3">
<Sparkles className="h-5 w-5" />
<h2 className="font-semibold">2) 見出しを指定して要約</h2>
</div>
<div className="flex flex-wrap gap-2 mb-2">
{presets.map(p => (
<button key={p} onClick={()=>setHeading(p)} className={`rounded-full border px-3 py-1 text-sm ${heading===p? 'bg-black text-white' : 'bg-white'}`}>{p}</button>
))}
</div>
<div className="flex flex-col sm:flex-row gap-3">
<input value={heading} onChange={e=>setHeading(e.target.value)} placeholder="例: Method" className="flex-1 rounded-lg border px-3 py-2" />
<input type="number" min={3} max={30} value={topK} onChange={e=>setTopK(Number(e.target.value))} className="w-28 rounded-lg border px-3 py-2" />
<button disabled={!canSummarize} onClick={doSummarize}
className="inline-flex items-center gap-2 rounded-lg bg-black px-4 py-2 text-white disabled:opacity-50">
{loading==='summarize' && <Loader2 className="h-4 w-4 animate-spin"/>}
Summarize
</button>
</div>
{error && <p className="mt-3 text-sm text-red-600">{error}</p>}
</section>


{summary && (
<section className="rounded-2xl border bg-white p-5 shadow-sm">
<h3 className="font-semibold mb-2">要約</h3>
<pre className="whitespace-pre-wrap text-sm leading-6">{summary}</pre>
<details className="mt-4">
<summary className="cursor-pointer text-sm text-gray-700">evidence を表示</summary>
<ul className="mt-2 space-y-2">
{evidence.map((ev, i) => (
<li key={i} className="rounded-lg border p-3 text-sm">
<div className="text-gray-500 mb-1">p.{ev.page ?? '-'} / {ev.section ?? ''} / {ev.heading ?? ''}</div>
<div className="text-gray-800">{String(ev.text ?? '').slice(0, 300)}{String(ev.text ?? '').length>300 ? '…' : ''}</div>
</li>
))}
</ul>
</details>
</section>
)}
</div>
)
}