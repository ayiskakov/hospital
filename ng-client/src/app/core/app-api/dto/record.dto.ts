import { CountryDto } from "./country.dto";
import { DiseaseDto } from "./disease.dto";
import { PublicServantDto } from "./public-servant.dto";

export interface RecordDto {
  email: string;
  cname: string;
  disease_code: string;
  total_deaths: number;
  total_patients: number;
  country: CountryDto;
  public_servant: PublicServantDto;
  disease: DiseaseDto;
}