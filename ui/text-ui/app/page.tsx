import { getMessages } from "@/apis/spam";
import ValidateSpam from "./components/ValidateSpam";
import MessageList from "./components/MessageList";

export default async function Home() {
  const messages =  await getMessages();
  console.log(messages)
  return (
    <main className="max-w-4xl mx-auto mt-4">
      <div className="text-center my-5 flex-col gap-4">
          <h1 className="text-2xl font-bold">Spam Message Detector</h1>
          <ValidateSpam />
      </div>

      <MessageList messages={messages} />
    </main>
  )
}
