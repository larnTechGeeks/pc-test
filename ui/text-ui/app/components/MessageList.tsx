import { IMessage } from '@/types/message'
import React from 'react'
import Message from './Message'

interface MessageListProps {
    messages: IMessage[]
}

const MessageList: React.FC<MessageListProps> = ({ messages }) => {
    return (
        <div className="overflow-x-auto">
            <table className="table">
                {/* head */}
                <thead>
                    <tr>
                        <th></th>
                        <th>Text</th>
                        <th>Is Spam</th>
                    </tr>
                </thead>
                <tbody>
                    {messages.map(message => (
                        <Message key={message.id} message={message} />
                    ))}
                </tbody>
            </table>
        </div>
    )
}

export default MessageList
