import UploadAndSummarize from '@/components/UploadAndSummarize'


export default function Page() {
return (
<div className="space-y-6">
<h1 className="text-2xl font-bold">論文要約エージェント</h1>
<p className="text-sm text-gray-600">
PDFをアップロードし、抽出されたチャンクをベクトルDBに登録した後、見出し名を指定して要約を取得します。
</p>
<UploadAndSummarize />
</div>
)
}