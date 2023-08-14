import { useMessageList } from "../hooks/useMessageList";

export const MessageList = () => {
  const messageList = useMessageList();

  return (
    <div>
      {messageList.map((m, i) => (
        <div key={i}>{m.content}</div>
      ))}
    </div>
  );
};