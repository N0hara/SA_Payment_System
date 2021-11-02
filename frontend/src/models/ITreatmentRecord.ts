import { AdmissionInterface } from "./IAdmission";

export interface TreatmentRecordInterface {
	ID: number, 
	Treatment: string,
	Cost: number,

	//RecordDate: Date | null,
    //FoodType    uint
	//MedAmount   uint
	//EquipmentID uint
	//MedID       string

	AdmissionID: number;
	Admission: AdmissionInterface;
}