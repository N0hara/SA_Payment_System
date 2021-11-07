import { RightTreatmentInterface } from "./IRightTreatment";
import { PatientInterface } from "./IPatient";

export interface AdmissionInterface {
    ID: number, 
    PatientID: number;
    Patient: PatientInterface;
    PatientName: string,
    RightTreatmentID: number;
	RightTreatment:   RightTreatmentInterface;
   
}