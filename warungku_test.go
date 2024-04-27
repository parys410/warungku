package warungku

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTambahBarang(t *testing.T) {
	myWarung := Warung{}
	myWarung.StrukBelanja = &[]Barang{}
	
	var myWarungIf WarungInterface = &myWarung

	barang1 := &Barang {
		Nama: "Cokelat",
		Harga: 3000,
		Qty: 5,
	}

	barang2 := &Barang {
		Nama: "Cokelat",
		Harga: 3000,
		Qty: 5,
	}

	myWarungIf.TambahBarang(barang1)
	myWarungIf.TambahBarang(barang2)

	assert.Equal(t, "Cokelat", (*myWarung.StrukBelanja)[0].Nama, "Harusnya Cokelat")
	assert.Equal(t, 10, (*myWarung.StrukBelanja)[0].Qty, "Qty Cokelat Harusnya 10")
	assert.Equal(t, 1, len((*myWarung.StrukBelanja)), "Jumlah Barang harusnya 1")
}

func TestKurangiQty(t *testing.T) {
	myWarung := Warung{}
	myWarung.StrukBelanja = &[]Barang{}
	
	var myWarungIf WarungInterface = &myWarung

	barang1 := &Barang {
		Nama: "Cokelat",
		Harga: 3000,
		Qty: 5,
	}

	barang2 := &Barang {
		Nama: "Pensil",
		Harga: 1000,
		Qty: 1,
	}

	myWarungIf.TambahBarang(barang1)
	myWarungIf.TambahBarang(barang2)

	myWarungIf.KurangiQty("Cokelat", 3)
	myWarungIf.KurangiQty("Pensil", 1)

	assert.Equal(t, 2, (*myWarung.StrukBelanja)[0].Qty, "Cokelat harusnya sisa 2")
	assert.Equal(t, 1, len((*myWarung.StrukBelanja)), "Jumlah Barang harusnya 1")
}