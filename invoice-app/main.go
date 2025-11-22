package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

type Item struct {
	Name  string
	Qty   int
	Price float64
	Total float64
}

type Invoice struct {
	StoreName   string
	UserName    string
	InvoiceNo   string
	Date        string
	Items       []Item
	TotalBill   float64
}

func main() {
	// Main page template with user input and modern UI
	tpl := template.Must(template.New("index").Parse(`
<!DOCTYPE html>
<html>
<head>
	<title>D-Mart Invoice System</title>
	<style>
		body { font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif; background: #f5f5f5; margin: 0; padding: 0; }
		.container { width: 80%; margin: 50px auto; background: #fff; padding: 30px; box-shadow: 0 0 20px rgba(0,0,0,0.1); border-radius: 10px; }
		h1 { text-align: center; color: #333; }
		label { font-weight: bold; }
		input[type=text], input[type=number] { width: 100%; padding: 8px; margin: 5px 0; border-radius: 5px; border: 1px solid #ccc; }
		table { width: 100%; border-collapse: collapse; margin-top: 20px; }
		th, td { padding: 12px; border-bottom: 1px solid #ddd; text-align: left; }
		th { background-color: #4CAF50; color: white; }
		tr:nth-child(even){background-color: #f9f9f9;}
		button { background-color: #4CAF50; color: white; border: none; padding: 12px 20px; margin-top: 20px; cursor: pointer; border-radius: 5px; font-size: 16px; transition: 0.3s; }
		button:hover { background-color: #45a049; }
	</style>
</head>
<body>
<div class="container">
	<h1>D-Mart Invoice System</h1>
	<form id="invoiceForm" target="_blank" method="POST" action="/generate">
		<label for="username">Customer Name:</label>
		<input type="text" name="username" placeholder="Enter customer name" required>

		<label for="invoiceno">Invoice Number (Optional):</label>
		<input type="text" name="invoiceno" placeholder="Auto-generated if empty">

		<table id="itemsTable">
			<tr>
				<th>Item Name</th>
				<th>Quantity</th>
				<th>Unit Price</th>
				<th>Total Price</th>
			</tr>
			<tr>
				<td><input type="text" name="name[]" required></td>
				<td><input type="number" name="qty[]" min="1" value="1" oninput="calculateRow(this)" required></td>
				<td><input type="number" name="price[]" min="0" step="0.01" value="0" oninput="calculateRow(this)" required></td>
				<td><input type="number" name="total[]" value="0" readonly></td>
			</tr>
		</table>
		<button type="button" onclick="addRow()">Add New Item</button>
		<br>
		<button type="submit">Generate Invoice</button>
	</form>
</div>

<script>
	function addRow() {
		const table = document.getElementById("itemsTable");
		const row = table.insertRow();
		row.innerHTML = '<td><input type="text" name="name[]" required></td>' +
		                '<td><input type="number" name="qty[]" min="1" value="1" oninput="calculateRow(this)" required></td>' +
		                '<td><input type="number" name="price[]" min="0" step="0.01" value="0" oninput="calculateRow(this)" required></td>' +
		                '<td><input type="number" name="total[]" value="0" readonly></td>';
	}

	function calculateRow(input) {
		const row = input.parentElement.parentElement;
		const qty = parseFloat(row.cells[1].children[0].value) || 0;
		const price = parseFloat(row.cells[2].children[0].value) || 0;
		row.cells[3].children[0].value = (qty * price).toFixed(2);
	}

	document.getElementById("invoiceForm").addEventListener("submit", function(e) {
		const totals = document.getElementsByName("total[]");
		let sum = 0;
		for (let t of totals) { sum += parseFloat(t.value) || 0; }
		const input = document.createElement("input");
		input.type = "hidden";
		input.name = "totalBill";
		input.value = sum.toFixed(2);
		this.appendChild(input);
	});
</script>
</body>
</html>
`))

	// Invoice template with improved UI
	invoiceTpl := template.Must(template.New("invoice").Parse(`
<!DOCTYPE html>
<html>
<head>
	<title>Invoice</title>
	<style>
		body { font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif; margin: 0; padding: 0; background: #f5f5f5; }
		.container { width: 80%; margin: 50px auto; background: #fff; padding: 30px; border-radius: 10px; box-shadow: 0 0 20px rgba(0,0,0,0.1); }
		h1, h3 { text-align: center; color: #333; }
		table { width: 100%; border-collapse: collapse; margin-top: 20px; }
		th, td { padding: 12px; border-bottom: 1px solid #ddd; text-align: left; }
		th { background-color: #4CAF50; color: white; }
		tr:nth-child(even){background-color: #f9f9f9;}
		#total { font-weight: bold; font-size: 18px; background-color: #e8f5e9; }
	</style>
</head>
<body>
<div class="container">
	<h1>{{.StoreName}}</h1>
	<h3>Customer: {{.UserName}}</h3>
	<h3>Invoice No: {{.InvoiceNo}} | Date: {{.Date}}</h3>

	<table>
		<tr>
			<th>Item Name</th>
			<th>Quantity</th>
			<th>Unit Price</th>
			<th>Total Price</th>
		</tr>
		{{range .Items}}
		<tr>
			<td>{{.Name}}</td>
			<td>{{.Qty}}</td>
			<td>{{printf "%.2f" .Price}}</td>
			<td>{{printf "%.2f" .Total}}</td>
		</tr>
		{{end}}
		<tr>
			<td colspan="3" id="total">Total Bill</td>
			<td id="total">{{printf "%.2f" .TotalBill}}</td>
		</tr>
	</table>
</div>
</body>
</html>
`))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	})

	http.HandleFunc("/generate", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		names := r.Form["name[]"]
		qtys := r.Form["qty[]"]
		prices := r.Form["price[]"]

		var items []Item
		for i := range names {
			var qty int
			var price float64
			fmt.Sscan(qtys[i], &qty)
			fmt.Sscan(prices[i], &price)
			items = append(items, Item{
				Name:  names[i],
				Qty:   qty,
				Price: price,
				Total: float64(qty) * price,
			})
		}

		var totalBill float64
		fmt.Sscan(r.FormValue("totalBill"), &totalBill)

		invoiceNo := r.FormValue("invoiceno")
		if invoiceNo == "" {
			invoiceNo = fmt.Sprintf("%d", time.Now().Unix())
		}

		invoice := Invoice{
			StoreName: "D-Mart Store",
			UserName:  r.FormValue("username"),
			InvoiceNo: invoiceNo,
			Date:      time.Now().Format("02-Jan-2006"),
			Items:     items,
			TotalBill: totalBill,
		}

		invoiceTpl.Execute(w, invoice)
	})

	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
