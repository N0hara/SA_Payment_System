import { useEffect, useState } from "react";
import { Link as RouterLink } from "react-router-dom";
import { createStyles, makeStyles, Theme } from "@material-ui/core/styles";
import Typography from "@material-ui/core/Typography";
import Button from "@material-ui/core/Button";
import Container from "@material-ui/core/Container";
import Paper from "@material-ui/core/Paper";
import Box from "@material-ui/core/Box";
import Table from "@material-ui/core/Table";
import TableBody from "@material-ui/core/TableBody";
import TableCell from "@material-ui/core/TableCell";
import TableContainer from "@material-ui/core/TableContainer";
import TableHead from "@material-ui/core/TableHead";
import TableRow from "@material-ui/core/TableRow";
import { BillInterface } from "../models/IBill";
import { format } from 'date-fns'

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    container: {
      marginTop: theme.spacing(2),
      minWidth: 1200,
    },
    table: {
      minWidth: 1100,
    },
    tableSpace: {
      marginTop: 20,
    },
    
  })
);

function Bills() {
  const classes = useStyles();
  const [bills, setBills] = useState<BillInterface[]>([]);
  const apiUrl = "http://localhost:8080";
  const requestOptions = {
    method: "GET",
    headers: {
      Authorization: `Bearer ${localStorage.getItem("token")}`,
      "Content-Type": "application/json",
    },
  };

  const getBills = async () => {
    fetch(`${apiUrl}/route/ListBill`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        console.log(res.data);
        if (res.data) {
            setBills(res.data);
        } else {
          console.log("else");
        }
      });
  };

  useEffect(() => {
    getBills();
  }, []);

  return (
    <div>
      <Container className={classes.container} maxWidth="md">
        <Box display="flex">
          <Box flexGrow={1}>
            <Typography
              component="h2"
              variant="h6"
              color="primary"
              style={{color:"black"}}
              gutterBottom
            >
              ???????????????????????????????????????????????????
            </Typography>
          </Box>
          <Box>
            <Button
              component={RouterLink}
              style={{backgroundColor:"#626567"}}
              to="/bill/create"
              variant="contained"
              color="primary"
            >
              ??????????????????????????????????????????????????????????????????
            </Button>
          </Box>
        </Box>
        <TableContainer component={Paper} className={classes.tableSpace}>
          <Table className={classes.table} aria-label="simple table">
            <TableHead>
              <TableRow>
                <TableCell align="center" width="5%">
                  ???????????????
                </TableCell>
                <TableCell align="center" width="15%">
                  ???????????????????????????????????????
                </TableCell>
                <TableCell align="center" width="10%">
                  ????????????????????????????????????????????????
                </TableCell>
                <TableCell align="center" width="10%">
                  ??????????????????????????????????????????????????????????????????????????????
                  </TableCell> 
                <TableCell align="center" width="10%">
                  ??????????????????????????????????????? ???????????????????????????
                </TableCell>
                <TableCell align="center" width="10%">
                  ??????????????????????????????????????????????????????
                </TableCell>
                <TableCell align="center" width="5%">
                  ???????????????????????????
                </TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {bills.map((item: BillInterface) => (
                <TableRow key={item.ID}>
                  <TableCell align="center">{item.ID}</TableCell>
                  <TableCell align="center">{format((new Date(item.BillDateTime)), 'dd MMMM yyyy hh:mm a')}</TableCell>
                  <TableCell align="center">{item.TreatmentRecordID}</TableCell>
                  <TableCell align="center">{item.TreatmentRecord.Admission.Patient.Identification}</TableCell>
                  <TableCell align="center">{item.RightTreatment.RightTreatmentName}</TableCell>
                  <TableCell align="center">{item.PaymentMethod.PaymentMethodName}</TableCell>
                  <TableCell align="center">{item.AmountPaid}</TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>
      </Container>
    </div>
  );
}

export default Bills;