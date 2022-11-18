import { CountryDto } from "./country.dto";
import { DiseaseDto } from "./disease.dto";

export interface DiscoveryDto {
  cname: string;
  disease_code: string;
  first_enc_date: string;
  country: CountryDto;
  disease: DiseaseDto;
}

