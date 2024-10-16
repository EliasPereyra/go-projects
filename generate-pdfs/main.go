package main

import (
	"log"
	"time"

	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/code"
	"github.com/johnfercher/maroto/v2/pkg/components/image"
	"github.com/johnfercher/maroto/v2/pkg/components/line"
	"github.com/johnfercher/maroto/v2/pkg/components/list"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/signature"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontfamily"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/consts/orientation"
	"github.com/johnfercher/maroto/v2/pkg/consts/pagesize"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

func main() {
	// We configure the basic format for the pdf. We set the view mode and view size
	cfg := config.NewBuilder().
		// In vertical view mode
		WithOrientation(orientation.Vertical).
		// this is the A4 format
		WithPageSize(pagesize.A4).
		WithLeftMargin(15).
		WithTopMargin(15).
		WithBottomMargin(15).
		WithRightMargin(15).
		Build()

	// We make an instance of maroto
	m := maroto.New(cfg)

	// Here we begin working with the sections of the pdf
	// 1. The Header
	addHeader(m)
	// 2. Invoice details
	addInvoiceDetails(m)
	// 3. Item List
	addItemList(m)
	// 4. Footer - signature and QR code
	addFooter(m)

	// Once we finish setting the sections, we generate the file
	document, err := m.Generate()
	if err != nil {
		log.Fatalf("There was an error when trying to generate the file: %v", err)
	}

	err = document.Save("output/invoice.pdf")
	if err != nil {
		log.Fatalf("There was an error when trying to save the file: %v", err)
	}

	log.Println("PDF saved successfully")
}

func addHeader(m core.Maroto) {
	m.AddRow(50,
		image.NewFromFileCol(12, "assets/go_logo.png",
			props.Rect{Center: true, Percent: 75}))

	m.AddRow(20, text.NewCol(12, "Elias Pereyra",
		props.Text{
			Top:   5,
			Style: fontstyle.Bold,
			Align: align.Center,
			Size:  16,
		}))

	m.AddRow(20, text.NewCol(12, "Factura", props.Text{
		Top:   5,
		Style: fontstyle.Bold,
		Size:  12,
		Align: align.Center,
	}))
}

func addInvoiceDetails(m core.Maroto) {
	m.AddRow(10,
		text.NewCol(6, "Date: "+time.Now().Format("02 Jan 2006"), props.Text{
			Align: align.Left,
			Size:  10,
			Style: fontstyle.Italic,
			Color: &props.Color{
				Red:   40,
				Green: 40,
				Blue:  40,
			},
		}),
		text.NewCol(6, "Invoice #1001", props.Text{
			Align: align.Right,
			Size:  10,
		}),
	)

	m.AddRow(10, line.NewCol(12))
}

type ItemList struct {
	Item            string
	Description     string
	Quantity        string
	Price           string
	DiscountedPrice string
	Total           string
}

// For generating a table we need to do two things:
// 1. Generate the header section
// 2. Generate the content section for each column

func (o ItemList) GetHeader() core.Row {
	return row.New(10).Add(
		text.NewCol(2, "Item", props.Text{Style: fontstyle.Bold}),
		text.NewCol(5, "Descripcion", props.Text{Style: fontstyle.Bold}),
		text.NewCol(1, "Cantidad", props.Text{Style: fontstyle.Bold}),
		text.NewCol(1, "Precio", props.Text{Style: fontstyle.Bold}),
		text.NewCol(2, "Descuento", props.Text{Style: fontstyle.Bold}),
		text.NewCol(1, "Total", props.Text{Style: fontstyle.Bold}),
	)
}

func (o ItemList) GetContent(i int) core.Row {
	r := row.New(5).Add(
		text.NewCol(2, o.Item, props.Text{
			Color: &props.Color{
				Red:   40,
				Green: 40,
				Blue:  40,
			},
		}),
		text.NewCol(5, o.Description, props.Text{
			Color: &props.Color{
				Red:   40,
				Green: 40,
				Blue:  40,
			},
		}),
		text.NewCol(1, o.Quantity, props.Text{
			Color: &props.Color{
				Red:   40,
				Green: 40,
				Blue:  40,
			},
		}),
		text.NewCol(1, o.Price, props.Text{
			Color: &props.Color{
				Red:   40,
				Green: 40,
				Blue:  40,
			},
		}),
		text.NewCol(2, o.DiscountedPrice, props.Text{
			Color: &props.Color{
				Red:   40,
				Green: 40,
				Blue:  40,
			},
		}),
		text.NewCol(1, o.Total, props.Text{
			Color: &props.Color{
				Red:   40,
				Green: 40,
				Blue:  40,
			},
		}),
	)

	if i%2 == 0 {
		r.WithStyle(&props.Cell{
			BackgroundColor: &props.Color{Red: 240, Green: 240, Blue: 240},
		})
	}

	return r
}

// We separate in another func to add objects to the table, as example
func getObjects() []ItemList {
	var list []ItemList

	contents := [][]string{
		{"Smartphone", "Teléfono móvil de última generación", "2", "$599.99", "$539.99", "$1079.98"},
		{"Tablet", "Computadora portátil de 10 pulgadas", "1", "$349.99", "$314.99", "$314.99", "$314.99"},
		{"Auriculares inalámbricos", "Audífonos Bluetooth con cancelación de ruido", "3", "$99.99", "$79.99", "$239.97", "$239.97"},
	}

	for i := 0; i < len(contents); i++ {
		list = append(list, ItemList{
			Item:            contents[i][0],
			Description:     contents[i][1],
			Quantity:        contents[i][2],
			Price:           contents[i][3],
			DiscountedPrice: contents[i][4],
			Total:           contents[i][5],
		})
	}

	return list
}

func addItemList(m core.Maroto) {
	rows, err := list.Build[ItemList](getObjects())
	if err != nil {
		log.Fatalf("There was an error when trying to build the list: %v", err.Error())
	}

	m.AddRows(rows...)
}

func addFooter(m core.Maroto) {
	m.AddRow(15, text.NewCol(10, "Total", props.Text{
		Top:   5,
		Style: fontstyle.Bold,
		Size:  12,
		Align: align.Right,
	}),
		text.NewCol(3, "$1100", props.Text{
			Top:   5,
			Style: fontstyle.Bold,
			Size:  10,
			Align: align.Center,
		}))

	m.AddRow(60, signature.NewCol(6, "Firma de Autorizacion", props.Signature{
		FontFamily: fontfamily.Courier,
	}),
		code.NewQrCol(10, "https://github.com/EliasPereyra", props.Rect{
			Percent: 75,
			Center:  true,
		}))
}
