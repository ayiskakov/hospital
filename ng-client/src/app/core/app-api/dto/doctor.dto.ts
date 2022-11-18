import { UserDto } from "./user.dto"
type Doctor = {
  degree: string;
}

export interface DoctorDto {
  user: UserDto;
  doctor: Doctor;
  degree?: string;
}
