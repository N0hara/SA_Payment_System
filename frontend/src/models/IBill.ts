import { RightTreatmentInterface } from "./IRightTreatment";
import { PaymentMethodInterface } from "./IPaymentMethod";
import { TreatmentRecordInterface } from "./ITreatmentRecord";

export interface BillInterface {
	ID: number, 
	BillDateTime: Date,
	AmountPaid: number,

	PatientPID: string;
	RightTreatmentID: number;
	RightTreatment:   RightTreatmentInterface;
	PaymentMethodID: number;
	PaymentMethod:   PaymentMethodInterface;
	TreatmentRecordID: number;
	TreatmentRecord:   TreatmentRecordInterface;

}