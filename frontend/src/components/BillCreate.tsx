import { useEffect, useState } from "react";
import { Link as RouterLink } from "react-router-dom";
import {
  makeStyles,
  Theme,
  createStyles,
  alpha,
} from "@material-ui/core/styles";
import Button from "@material-ui/core/Button";
import FormControl from "@material-ui/core/FormControl";
import Container from "@material-ui/core/Container";
import Paper from "@material-ui/core/Paper";
import Grid from "@material-ui/core/Grid";
import Box from "@material-ui/core/Box";
import Typography from "@material-ui/core/Typography";
import Divider from "@material-ui/core/Divider";
import Snackbar from "@material-ui/core/Snackbar";
import Select from "@material-ui/core/Select";
import MuiAlert, { AlertProps } from "@material-ui/lab/Alert";

import { RightTreatmentInterface } from "../models/IRightTreatment";
import { PaymentMethodInterface } from "../models/IPaymentMethod";
import { PatientInterface } from "../models/IPatient";
import { AdmissionInterface } from "../models/IAdmission";
import { TreatmentRecordInterface } from "../models/ITreatmentRecord"; 
import { BillInterface } from "../models/IBill";
import { TextField } from "@material-ui/core";

import {
  MuiPickersUtilsProvider,
  KeyboardDateTimePicker,
} from "@material-ui/pickers";
import DateFnsUtils from "@date-io/date-fns";

const Alert = (props: AlertProps) => {
  return <MuiAlert elevation={6} variant="filled" {...props} />;
};

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      flexGrow: 1,
    },
    container: {
      marginTop: theme.spacing(2),
    },
    paper: {
      padding: theme.spacing(2),
      color: theme.palette.text.secondary,
    },
  })
);

function BillCreate() {
  const classes = useStyles();
  const [selectedDate, setSelectedDate] = useState<Date | null>(new Date());
  const [righttreatments, setRightTreatments] = useState<RightTreatmentInterface[]>([]);
  const [paymentmethods, setPaymentMethods] = useState<PaymentMethodInterface[]>([]);
  const [patients, setPatients] = useState<PatientInterface[]>([]);
  const [admissions, setAdmissions] = useState<AdmissionInterface[]>([]);
  const [treatmentrecords, setTreatmentRecords] = useState<TreatmentRecordInterface[]>([]);
  const [bill, setBill] = useState<Partial<BillInterface>>(
    {}
  );

  const [success, setSuccess] = useState(false);
  const [error, setError] = useState(false);

  const apiUrl = "http://localhost:8080";
  const requestOptions = {
    method: "GET",
    headers: {
      Authorization: `Bearer ${localStorage.getItem("token")}`,
      "Content-Type": "application/json",
    },
  };

  const handleInputChange = (

    event: React.ChangeEvent<{ id?: string; value: any }>
 
  ) => {
 
    const id = event.target.id as keyof typeof BillCreate;
 
    const { value } = event.target;
 
    setBill({ ...bill, [id]: value });
 

  };
  const handleClose = (event?: React.SyntheticEvent, reason?: string) => {
    if (reason === "clickaway") {
      return;
    }
    setSuccess(false);
    setError(false);
  };

  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: unknown }>
  ) => {
    const name = event.target.name as keyof typeof bill;
    setBill({
      ...bill,
      [name]: event.target.value,
    });
  };

  const handleDateChange = (date: Date | null) => {
    console.log(date);
    setSelectedDate(date);
  };
//=================================================================================================================================
//=================================================================================================================================
//=================================================================================================================================

  /*
  const getTreatment = async () => {
    var tmpTreat =treatments;
    fetch(`${apiUrl}/TreatmentRecords`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          console.log("Here is the treatment data");
          console.log(res.data);
          //setTreatments(res.data);
          tmpTreat = res.data;
        } else {
          console.log("else");
        }
      });
      fetch(`${apiUrl}/admissions`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if(res.data){
          console.log("here is a tmptreat");
          console.log(tmpTreat);
          var newTmp = tmpTreat.map((item,index) => {
            item.Admission = res.data[index];
            return item;
          });
          console.log("here is a newtmp");
          console.log(newTmp);
          setTreatments(newTmp);
        }else{
            console.log("error admission");
        }
      })
  };*/
  
  const getRightTreatment = async () => {
    fetch(`${apiUrl}/righttreatments`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setRightTreatments (res.data);
        } else {
          console.log("else");
        }
      });
  };
  const getPaymentMethod = async () => {
    fetch(`${apiUrl}/paymentmethods`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
            setPaymentMethods(res.data);
        } else {
          console.log("else");
        }
      });
  };
  const getPatient = async () => {
    fetch(`${apiUrl}/patients`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
            setPatients(res.data);
        } else {
          console.log("else");
        }
      });
  };

  const getAdmission= async () => {
    fetch(`${apiUrl}/patient/admissions`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setAdmissions(res.data);
        } else {
          console.log("else");
        }
      });
  };
  
  const getTreatmentRecord= async () => {
    fetch(`${apiUrl}/admission/treatmentrecords`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setTreatmentRecords(res.data);
        } else {
          console.log("else");
        }
      });
  };

  


/*
  const getAdmission = async () => {
    let uid = localStorage.getItem("uid");
    fetch(`${apiUrl}/playlist/watched/user/${uid}`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        watchVideo.PlaylistID = res.data.ID
        if (res.data) {
          setPlaylists(res.data);
        } else {
          console.log("else");
        }
      });
  };*/

  useEffect(() => {
    getRightTreatment();
    getPaymentMethod();
    getPatient();
    getAdmission();
    getTreatmentRecord();
  }, []);



  const convertType = (data: string | number | undefined) => {
    let val = typeof data === "string" ? parseInt(data) : data;
    return val;
  };
 
  function submit() {
    let data = {
        BillDateTime: selectedDate,
        AmountPaid: convertType(bill.TreatmentRecord?.Cost),
        IdentificationID: convertType(bill.TreatmentRecord?.Admission.Patient.IdentificationID),
        //PatientID: convertType(bill.PatientID),
        RightTreatmentID:  convertType(bill.RightTreatmentID),
        PaymentMethodID: convertType(bill.PaymentMethodID),
        TreatmentRecordID: convertType(bill.TreatmentRecordID),
        
    };

    console.log(data)

    const requestOptionsPost = {
      method: "POST",
      headers: {
        Authorization: `Bearer ${localStorage.getItem("token")}`,
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    };

    fetch(`${apiUrl}/bills`, requestOptionsPost)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          console.log("บันทึกได้")
          setSuccess(true);
        } else {
          console.log("บันทึกไม่ได้")
          setError(true);
        }
      });
  }

  return (
    <Container className={classes.container} maxWidth="md">
      <Snackbar open={success} autoHideDuration={6000} onClose={handleClose}>
        <Alert onClose={handleClose} severity="success">
          บันทึกข้อมูลสำเร็จ
        </Alert>
      </Snackbar>
      <Snackbar open={error} autoHideDuration={6000} onClose={handleClose}>
        <Alert onClose={handleClose} severity="error">
          บันทึกข้อมูลไม่สำเร็จ
        </Alert>
      </Snackbar>
      <Paper className={classes.paper}>
        <Box display="flex">
          <Box flexGrow={1}>
            <Typography
              component="h2"
              variant="h6"
              color="primary"
              gutterBottom
            >
              บันทึกการชำระเงิน
            </Typography>
          </Box>
        </Box>
        <Divider />
        <Grid container spacing={3} className={classes.root}>
        {/*
        <Grid item xs={6}>

            //คอมเมนต์============================================
            <FormControl fullWidth variant="outlined">
              <p>ผู้จ่ายยา</p>
              <Select
                native
                disabled
                value={medRecord.PharmaID}
               // onChange={handleChange}
                /*inputProps={{
                  name: "PharmaID",
                }}
              >
                <option aria-label="None" value="">

                  {pharmacists?.Name}
                </option>
              </Select>
            </FormControl>
            //คอมเมนต์============================================

          </Grid>
                */}
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>ใบบันทึกการรักษาหมายเลข</p>
              <Select
                native
                value={bill.TreatmentRecordID}
                onChange={handleChange}
                inputProps={{
                  name: "TreatmentRecordID",
                }}
              >
                <option aria-label="None" value="">
                  กรุณาเลือกใบบันทึกการรักษา
                </option>
               
                {treatmentrecords.map((item: TreatmentRecordInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.ID} {item.Admission.PatientName}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>สิทธิการรักษา หรือประกัน</p>
              <Select
                native
                value={bill.RightTreatmentID}
                onChange ={handleChange}
                inputProps = {{
                  name: "RightTreatmentID",
                }}
              >
                <option aria-label="None" value="">
                    กรุณาเลือกสิทธิการรักษา หรือประกัน
                </option>
                {righttreatments.map((item: RightTreatmentInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.RightTreatmentName}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>

          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>ช่องทางการชำระเงิน</p>
              <Select
                native
                value={bill.PaymentMethodID}
                onChange={handleChange}
                inputProps={{
                  name: "PaymentMethodID",
                }}
              >
                <option aria-label="None" value="">
                    กรุณาเลือกช่องทางการชำระเงิน
                </option>
                {paymentmethods.map((item: PaymentMethodInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.PaymentMethodName}
                  </option>
                ))}
              </Select>
            </FormControl>
          </Grid>
        {/*
        //คอมเมนต์ช=====================================================
          <Grid item xs={6}>
          <FormControl fullWidth variant="outlined">
            <p>จำนวน</p>
            <TextField
              label="กรุณากรอกจำนวน"
              id="Amount"
              name="Amount"
              variant="outlined"
              type="string"
              size="medium"
              value={medRecord.Amount || ""}
              onChange={handleInputChange}
            />
          </FormControl>
          </Grid>
        */}
          
          {
          //คอมเมนต์ช=====================================================
          /*
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>test</p>
              <Select
                native
                value={medRecord.AdmissionID}
                onChange={handleChange}
                inputProps={{
                  name: "AdmissionID",
                }}
              >
                <option aria-label="None" value="">
                  กรุณา test
                </option>
                {admissions.map((item: AdmissionInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.PatientID}  {item.Patientname}
                  </option>
                ))}
              </Select>
            </FormControl>
                </Grid>*/}
          {/*    field lock
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>เพลย์ลิสต์</p>
              <Select
                native
                value={medRecord.AdmissionID}
                onChange={handleChange}
                disabled
                inputProps={{
                  name: "AdmissionID",
                }}
              >
                <option aria-label="None" value="">
                  กรุณาเลือกเพลย์ลิสต์
                </option>
                <option value={medRecord?.ID} key={medRecord?.ID}>
                  {medRecord?.Admission?.Patientname}
                </option>
                
              </Select>
            </FormControl>
              </Grid> */}
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined">
              <p>วันที่และเวลา</p>
              <MuiPickersUtilsProvider utils={DateFnsUtils}>
                <KeyboardDateTimePicker
                  name="Billdatetime"
                  value={selectedDate}
                  onChange={handleDateChange}
                  label=""
                  minDate={new Date("2018-01-01T00:00")}
                  format="yyyy/MM/dd hh:mm a"
                />
              </MuiPickersUtilsProvider>
            </FormControl>
          </Grid>
          <Grid item xs={12}>
            <Button
              component={RouterLink}
              to="/bills"
              variant="contained"
            >
              กลับ
            </Button>
            <Button
              style={{ float: "right" ,backgroundColor:"#626567" }}
              variant="contained"
              onClick={submit}
              color="primary"
            >
              บันทึก
            </Button>
          </Grid>
        </Grid>
      </Paper>
    </Container>
  );
}

export default BillCreate;