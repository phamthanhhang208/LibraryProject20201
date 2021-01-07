package doc

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/pinezapple/LibraryProject20201/portal/controller"
	portalModel "github.com/pinezapple/LibraryProject20201/portal/model"
)

func SelectAllAvailableBarcode(c echo.Context) (erro error) {
	return controller.ExecHandler(c, nil, selectAllBarcodeAvail)
}

func SelectAllSellingBarcode(c echo.Context) (erro error) {
	return controller.ExecHandler(c, nil, selectAllBarcodeSelling)
}

func SelectAllDamagedBarcode(c echo.Context) (erro error) {
	return controller.ExecHandler(c, nil, selectAllBarcodeDamaged)
}

func SelectAllBorrowForm(c echo.Context) (erro error) {
	return controller.ExecHandler(c, nil, selectAllBorrowForm)
}

func SelectAllNotReturnedBorrowForm(c echo.Context) (erro error) {
	return controller.ExecHandler(c, nil, selectAllUnreturnedBorrowForm)
}

func SelectAllPayment(c echo.Context) (erro error) {
	return controller.ExecHandler(c, nil, selectAllPayment)
}

func SelectAllSaleBill(c echo.Context) (erro error) {
	return controller.ExecHandler(c, nil, selectAllSaleBill)
}

func SelectBarcodeByID(c echo.Context) (erro error) {
	return controller.ExecHandler(c, &portalModel.SelectBarcodeByIDReq{}, selectBarcodeByID)
}

func SelectBorrowFormByID(c echo.Context) (erro error) {
	return controller.ExecHandler(c, &portalModel.SelectBorrowFormByIDReq{}, selectBorrowFormByID)
}

func SelectPaymentByID(c echo.Context) (erro error) {
	return controller.ExecHandler(c, &portalModel.SelectPaymentByIDReq{}, selectPaymentByID)
}

func SelectSaleBillByID(c echo.Context) (erro error) {
	return controller.ExecHandler(c, &portalModel.SelectSaleBillByIDReq{}, selectSaleBillByID)
}

func SaveDocumentByBatch(c echo.Context) (erro error) {
	fmt.Println("In Save Document By Batch")
	return
}
