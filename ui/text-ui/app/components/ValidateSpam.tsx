"use client"

import { CiCirclePlus } from "react-icons/ci";
import { FormEventHandler, useState } from "react";
import { IMessageRequest } from "@/types/message";
import { classifyMessage } from "@/apis/spam";

const ValidateSpam = () => {

  const [newMessage, setNewMessage] = useState<string>('')

  const handleSubmitForm: FormEventHandler<HTMLFormElement> = async (e) => {
    e.preventDefault();

    const message:IMessageRequest = {
      text: newMessage,
    }

    await classifyMessage(message);

    setNewMessage("");
  }

  return (
    <div>

      <form className="space-y-4" onSubmit={handleSubmitForm}>
        <div>
          <label className="label">
            <span className="text-base label-text">Message</span>
          </label>
          <input value={newMessage} onChange={e => setNewMessage(e.target.value)} type="text" placeholder="Message" className="w-full input input-bordered input-primary" />
        </div>
        <div>
          <button className='btn btn-primary w-full'>Validate Spam
            <CiCirclePlus size={21} /></button>
        </div>
      </form>
    </div>
  )

}

export default ValidateSpam
