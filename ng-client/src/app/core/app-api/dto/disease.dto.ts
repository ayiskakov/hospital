export interface DiseaseDto {
  id: number;
  description: string;
  disease_code: string;
  pathogen: string;
  disease_type: { id: number, description: string };
}