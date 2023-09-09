import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgRegister } from "./types/eywacontract/eywacontract/tx";
import { MsgHandshake } from "./types/eywacontract/eywacontract/tx";
import { MsgSendChat } from "./types/eywacontract/eywacontract/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/eywacontract.eywacontract.MsgRegister", MsgRegister],
    ["/eywacontract.eywacontract.MsgHandshake", MsgHandshake],
    ["/eywacontract.eywacontract.MsgSendChat", MsgSendChat],
    
];

export { msgTypes }