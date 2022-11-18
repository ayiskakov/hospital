import { DiseaseTypeDto } from "./disease-type.dto";
import { DoctorDto } from "./doctor.dto";

export interface SpecializeDto {
    id: number;
    email: string;
    disease_type: DiseaseTypeDto;
    doctor: DoctorDto;
  }