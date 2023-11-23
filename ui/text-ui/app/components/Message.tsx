import { IMessage } from '@/types/message'
import React, { useState } from 'react'

interface MessageProps {
    message: IMessage
}

const Message:React.FC<MessageProps> = ({ message }) => {
    return (
        <tr key={message.id} className="bg-base-200">
            <th>{message.id}</th>
            <td>{message.text}</td>
            <td>{message.spam}</td>
        </tr>
    )
}

export default Message
