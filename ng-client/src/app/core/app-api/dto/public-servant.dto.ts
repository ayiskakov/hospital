import { UserDto } from "./user.dto"

type PublicServant = {
    department: string;
}

export interface PublicServantDto {
    user: UserDto;
    public_servant: PublicServant;
    department?: string;
}