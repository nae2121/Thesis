import './globals.css'
import type { Metadata } from 'next'


export const metadata: Metadata = {
title: 'PaperAgent',
description: 'PDF ingest & RAG summary UI',
}


export default function RootLayout({ children }: { children: React.ReactNode }) {
return (
<html lang="ja">
<body>
<div className="min-h-screen">
<header className="border-b bg-white">
<div className="mx-auto max-w-4xl px-4 py-3 flex items-center gap-2">
<span className="font-semibold">PaperAgent</span>
<span className="text-xs text-gray-500">Next.js UI</span>
</div>
</header>
<main className="mx-auto max-w-4xl px-4 py-8">{children}</main>
</div>
</body>
</html>
)
}