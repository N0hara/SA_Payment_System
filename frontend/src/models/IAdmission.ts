import { PatientInterface } from "./IPatient";

export interface AdmissionInterface {
    ID: number, 
    PatientID: number;
    Patient: PatientInterface;
    PatientName: string,
   
}