import { createStyles, makeStyles, Theme } from "@material-ui/core/styles";
import Container from "@material-ui/core/Container";

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    container: {
      marginTop: theme.spacing(2),
    },
    table: {
      minWidth: 650,
    },
    tableSpace: {
      marginTop: 20,
    },
  })
);

function Home() {
  const classes = useStyles();

  return (
    <div>
      <Container className={classes.container} maxWidth="md">
        <h1 style={{ textAlign: "center" }}>ระบบบันทึกการชำระเงินผู้ป่วยใน</h1>
        <h4>Requirements</h4>
        <p>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
        ระบบการชำระเงิน ภายในโรงพยาบาล เป็นระบบที่ให้เจ้าหน้าที่การเงิน กรอกเลขที่ใบบันทึกการรักษา จากนั้นดึงข้อมูลค่าใช้จ่ายของผู้ป่วยในจากระบบบันทึกการ เพื่อให้เจ้าหน้าที่การเงินได้ดำเนินการชำระเงินต่อไป รวมถึงการเลือกสิทธิค่ารักษาพยาบาล หรือประกันต่างๆของผู้ป่วยใน เพื่อให้ได้จำนวนเงินในการนำไปคำนวณ และเลือกช่องทางการชำระเงิน โดยเมื่อชำระเงินเสร็จแล้วก็ทำการบันทึกการชำระเงินนั้น และจะได้ใบบันทึกการชำระเงิน
        </p>
        <p>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
        นอกจากความสามารถเหล่านี้แล้ว หากเจ้าหน้าที่การเงิน ต้องการแก้ไขหรือ เพิ่มเติมข้อมูลของตน ต้องติดต่อผู้ดูแลระบบเท่านั้น ไม่สามารถแก้ไขใดๆเองได้ แต่เจ้าหน้าที่การเงินสามารถแก้ไขส่วนของสิทธิค่ารักษาพยาบาล หรือประกันต่างๆของผู้ป่วยในได้
        </p>
        <img src="/img/HomePage.jpg" width="900px"></img>
      </Container>
    </div>
  );
}
export default Home;