import { Injectable } from '@angular/core';
import { HttpClient } from "@angular/common/http";
import { Observable } from "rxjs";
import { CountryDto, DoctorDto, UserDto, PublicServantDto, DiseaseTypeDto, DiseaseDto, DiscoveryDto, RecordDto, SpecializeDto } from "./dto";
import { AppApiModule } from "./app-api.module";
import { environment } from "../../../environments/environment";

const PATH = environment.api;

@Injectable({
  providedIn: AppApiModule
})

export class AppApiService {

  constructor(private http: HttpClient) { }

  // Country

  getCountryList(): Observable<CountryDto[]> {
    return this.http.get<CountryDto[]>(PATH + "countries")
  }

  createCountry(country: CountryDto): Observable<CountryDto> {
    return this.http.post<CountryDto>(PATH + 'countries', country)
  }

  updateCountry(cname: string, country: CountryDto): Observable<CountryDto> {
    return this.http.put<CountryDto>(PATH + 'countries/' + cname, country)
  }

  deleteCountry(cname: string): Observable<CountryDto> {
    return this.http.delete<CountryDto>(PATH + 'countries/' + cname)
  }

  // User

  getUserList(): Observable<UserDto[]> {
    return this.http.get<UserDto[]>(PATH + "users")
  }

  createUser(user: { user: UserDto }): Observable<UserDto> {
    return this.http.post<UserDto>(PATH + 'users', user)
  }

  updateUser(email: string, user: UserDto): Observable<UserDto> {
    return this.http.put<UserDto>(PATH + 'users/' + email, user)
  }

  deleteUser(email: string): Observable<UserDto> {
    return this.http.delete<UserDto>(PATH + 'users/' + email)
  }

  // Doctor

  getDoctorList(): Observable<DoctorDto[]> {
    return this.http.get<DoctorDto[]>(PATH + "doctors")
  }

  createDoctor(doctor: DoctorDto): Observable<DoctorDto> {
    return this.http.post<DoctorDto>(PATH + 'doctors', doctor)
  }

  updateDoctor(email: string, doctor: UserDto): Observable<DoctorDto> {
    return this.http.put<DoctorDto>(PATH + 'doctors/' + email, doctor)
  }

  deleteDoctor(email: string): Observable<{ email: string }> {
    return this.http.delete<{ email: string }>(PATH + 'doctors/' + email)
  }

  // Public Servant

  getPublicServantList(): Observable<PublicServantDto[]> {
    return this.http.get<PublicServantDto[]>(PATH + "publicServants")
  }

  createPublicServant(publicServant: PublicServantDto): Observable<PublicServantDto> {
    return this.http.post<PublicServantDto>(PATH + 'publicServants', publicServant)
  }

  updatePublicServant(email: string, publicServant: PublicServantDto): Observable<PublicServantDto> {
    return this.http.put<PublicServantDto>(PATH + 'publicServants/' + email, publicServant)
  }

  deletePublicServant(email: string): Observable<PublicServantDto> {
    return this.http.delete<PublicServantDto>(PATH + 'publicServants/' + email)
  }

  // Disease Type

  getDiseaseTypeList(): Observable<DiseaseTypeDto[]> {
    return this.http.get<DiseaseTypeDto[]>(PATH + "diseaseTypes")
  }

  createDiseaseType(diseaseType: DiseaseTypeDto): Observable<DiseaseTypeDto> {
    return this.http.post<DiseaseTypeDto>(PATH + 'diseaseTypes', diseaseType)
  }

  updateDiseaseType(id: number, diseaseType: DiseaseTypeDto): Observable<DiseaseTypeDto> {
    return this.http.put<DiseaseTypeDto>(PATH + 'diseaseTypes/' + id, diseaseType)
  }

  deleteDiseaseType(id: number): Observable<DiseaseTypeDto> {
    return this.http.delete<DiseaseTypeDto>(PATH + 'diseaseTypes/' + id)
  }

  // Disease

  getDiseaseList(): Observable<DiseaseDto[]> {
    return this.http.get<DiseaseDto[]>(PATH + "diseases")
  }

  createDisease(disease: DiseaseDto): Observable<DiseaseDto> {
    return this.http.post<DiseaseDto>(PATH + 'diseases', disease)
  }

  updateDisease(disease_code: string, disease: DiseaseDto): Observable<DiseaseDto> {
    return this.http.put<DiseaseDto>(PATH + 'diseases/' + disease_code, disease)
  }

  deleteDisease(disease_code: string): Observable<DiseaseDto> {
    return this.http.delete<DiseaseDto>(PATH + 'diseases/' + disease_code)
  }

  // Discovery

  getDiscoveryList(): Observable<DiscoveryDto[]> {
    return this.http.get<DiscoveryDto[]>(PATH + "discovery")
  }

  createDiscovery(discovery: DiscoveryDto): Observable<DiscoveryDto> {
    return this.http.post<DiscoveryDto>(PATH + 'discovery', discovery)
  }

  updateDiscovery(cname: string, disease_code: string, discovery: DiscoveryDto): Observable<DiscoveryDto> {
    return this.http.put<DiscoveryDto>(PATH + 'discovery/' + cname + "/" + disease_code, discovery)
  }

  deleteDiscovery(cname: string, disease_code: string): Observable<DiscoveryDto> {
    return this.http.delete<DiscoveryDto>(PATH + 'discovery/' + cname + "/" + disease_code)
  }

  // Records

  getRecordList(): Observable<RecordDto[]> {
    return this.http.get<RecordDto[]>(PATH + "records")
  }

  createRecord(record: RecordDto): Observable<RecordDto> {
    return this.http.post<RecordDto>(PATH + 'records', record)
  }

  updateRecord(email: string, cname: string, disease_code: string, record: { total_deaths: number, total_patients: number }): Observable<RecordDto> {
    return this.http.put<RecordDto>(PATH + 'records/' + email + "/" + cname + "/" + disease_code, record)
  }

  deleteRecord(email: string, cname: string, disease_code: string): Observable<RecordDto> {
    return this.http.delete<RecordDto>(PATH + 'records/' + email + "/" + cname + "/" + disease_code)
  }

  // Specialize

  getSpecializeList(): Observable<SpecializeDto[]> {
    return this.http.get<SpecializeDto[]>(PATH + "specialize")
  }

  createSpecialize(specialize: SpecializeDto): Observable<SpecializeDto> {
    return this.http.post<SpecializeDto>(PATH + 'specialize', specialize)
  }

  deleteSpecialize(email: string, id: number): Observable<SpecializeDto> {
    return this.http.delete<SpecializeDto>(PATH + 'specialize/' + email + "/" + id)
  }
}
