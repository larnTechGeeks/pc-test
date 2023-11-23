import { IMessage, IMessageRequest } from "../types/message"

const baseURL = 'http://localhost:9000/v1'

export const getMessages = async(): Promise<IMessage[]> => {
    const res = await fetch(`${baseURL}/messages`, { cache: 'no-store' });
    const messages = await res.json();
    console.log(messages);
    return messages;
}


export const classifyMessage = async(message: IMessageRequest): Promise<IMessage> => {
    const res = await fetch(`${baseURL}/messages`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },

        body: JSON.stringify(message),
    });

    const classifiedMessage = await res.json();
    return classifiedMessage;
}