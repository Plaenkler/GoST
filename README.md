# 👻 G(h)oST

is a *small* project whose core component is the software **STrack**. The application serves the purpose of visualizing key data of invoices from the [Office Management Systems](https://oms.hctoms.de/) module Sales. In this way, viewers can gain valuable insights for managing their business. 

> The name **G(h)oST** is a tribute to Golang and a small blow against the obsolete OMS.

<table style="border:none;">
  <tr>
    <td><img src="https://user-images.githubusercontent.com/60503970/175663806-7e8338c2-d608-468c-a799-fefe89ed7261.png" width="480"/></td>
    <td><img src="https://user-images.githubusercontent.com/60503970/175665267-bd11a17d-f9f5-4d91-a6bc-506d1dcf994f.png" width="480"/></td>
  </tr>
</table>

## ⚙️ How it works

Currently, the software creates an empty database that must be filled separately with data from the OMS. A small web server then provides the data in processed form in simple charts and cards.

## 🎯 Project goals

- [x] Persistent database for data copies
- [x] Data preparation for quantity relations
- [x] Provide data in web frontend
- [ ] Independent data import process
- [ ] Detailed display for invoices and products

## 👪 Attribution


- [Valentin Kaiser](https://github.com/Valentin-Kaiser) provides the go module [go-dbase](https://pkg.go.dev/github.com/Valentin-Kaiser/go-dbase), which allows the import of data directly via STrack.
- [Waschbaer](https://github.com/daWaschbaer) provided a way to import data and helped me analyze the OMS data structures.

## 📜 Installation guide





