import { LabeledText } from '@/components/form/labeledText'
import { PageTitle } from '@/components/titles'
import { getUserById } from '@/lib/api/user'

export default async function UserPage({ params }: { params: { id: number } }) {
  const user = await getUserById(params.id)

  return (
    <div>
      <PageTitle title={user.name} />
      <div className="flex flex-col gap-y-2">
        <LabeledText label="ID" value={user.id.toString()} isCode={true} />
        <LabeledText label="Name" value={user.name} />
        <LabeledText label="Motto" value={user.motto} />
        <LabeledText label="Pet Name" value={user.petName || '-'} />
        <LabeledText label="Favorite Food" value={user.favoriteFood} />
      </div>
    </div>
  )
}
