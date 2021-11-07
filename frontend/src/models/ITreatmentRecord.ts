import { AdmissionInterface } from "./IAdmission";

export interface TreatmentRecordInterface {
	ID: number, 
	Treatment: string,
	Cost: number,

	AdmissionID: number;
	Admission: AdmissionInterface;
}