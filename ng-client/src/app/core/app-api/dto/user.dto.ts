import { CountryDto } from "./country.dto";

export interface UserDto {
  name: string;
  surname: string;
  email: string;
  phone: string;
  salary: string;
  country: CountryDto;
}