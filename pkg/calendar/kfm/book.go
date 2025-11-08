package kfm

import (
	"strings"

	"github.com/pkg/errors"
)

type Book int

const (
	Genesis Book = iota
	Exodus
	Levitikus
	Numeri
	Deuteronomium
	Josua
	Richter
	Rut
	ErsteSamuel
	ZweiteSamuel
	ErsteKoenige
	ZweiteKoenige
	ErsteChronik
	ZweiteChronik
	Esra
	Nehemia
	Ester
	Ijob
	Psalmen
	Sprichwoerter
	Kohelet
	Hohelied
	Jesaja
	Jeremia
	Klagelieder
	Ezechiel
	Daniel
	Hosea
	Joel
	Amos
	Obadja
	Jona
	Micha
	Nahum
	Habakuk
	Zefanja
	Haggai
	Sacharja
	Maleachi
	Mattaeus
	Markus
	Lukas
	Johannes
	Apostelgeschichte
	Roemer
	ErsteKorinther
	ZweiteKorinther
	Galater
	Epheser
	Philipper
	Kolosser
	ErsteThessalonicher
	ZweiteThessalonicher
	ErsteTimotheus
	ZweiteTimotheus
	Titus
	Philemon
	Hebraer
	Jakobus
	ErstePetrus
	ZweitePetrus
	ErsteJohannes
	ZweiteJohannes
	DritteJohannes
	Judas
	Offenbarung
	ErsteNephi
	ZweiteNephi
	Jakob
	Enos
	Jarom
	Omni
	WorteMormons
	Mosia
	Alma
	Heleman
	DritteNephi
	VierteNephi
	Mormon
	Ether
	Moroni
	LehreUndBuendnisse
	AmtlicheErklaerung1
	AmtlicheErklaerung2
	Mose
	Abraham
	JosephSmithMatthaeus
	JosephSmithLebensgeschichte
	Glaubensartikel
)

func ParseBook(s string) (Book, error) {
	switch strings.ToLower(s) {
	case "gen":
		return Genesis, nil
	case "ex":
		return Exodus, nil
	case "lev":
		return Levitikus, nil
	case "num":
		return Numeri, nil
	case "dtn":
		return Deuteronomium, nil
	case "jos":
		return Josua, nil
	case "ri":
		return Richter, nil
	case "rut":
		return Rut, nil
	case "1 sam":
		return ErsteSamuel, nil
	case "2 sam":
		return ZweiteSamuel, nil
	case "1 kön":
		return ErsteKoenige, nil
	case "2 kön":
		return ZweiteKoenige, nil
	case "1 chr":
		return ErsteChronik, nil
	case "2 chr":
		return ZweiteChronik, nil
	case "esra":
		return Esra, nil
	case "neh":
		return Nehemia, nil
	case "est":
		return Ester, nil
	case "ijob":
		return Ijob, nil
	case "ps":
		return Psalmen, nil
	case "spr":
		return Sprichwoerter, nil
	case "koh":
		return Kohelet, nil
	case "hld":
		return Hohelied, nil
	case "jes":
		return Jesaja, nil
	case "jer":
		return Jeremia, nil
	case "klgl":
		return Klagelieder, nil
	case "ez":
		return Ezechiel, nil
	case "dan":
		return Daniel, nil
	case "hos":
		return Hosea, nil
	case "joël":
		return Joel, nil
	case "am":
		return Amos, nil
	case "obd":
		return Obadja, nil
	case "jona":
		return Jona, nil
	case "mi":
		return Micha, nil
	case "nah":
		return Nahum, nil
	case "hab":
		return Habakuk, nil
	case "zef":
		return Zefanja, nil
	case "hag":
		return Haggai, nil
	case "sach":
		return Sacharja, nil
	case "mal":
		return Maleachi, nil
	case "mt":
		return Mattaeus, nil
	case "mk":
		return Markus, nil
	case "lk":
		return Lukas, nil
	case "joh":
		return Johannes, nil
	case "apg":
		return Apostelgeschichte, nil
	case "röm":
		return Roemer, nil
	case "1 kor":
		return ErsteKorinther, nil
	case "2 kor":
		return ZweiteKorinther, nil
	case "gal":
		return Galater, nil
	case "eph":
		return Epheser, nil
	case "phil":
		return Philipper, nil
	case "kol":
		return Kolosser, nil
	case "1 thess":
		return ErsteThessalonicher, nil
	case "2 thess":
		return ZweiteThessalonicher, nil
	case "1 tim":
		return ErsteTimotheus, nil
	case "2 tim":
		return ZweiteTimotheus, nil
	case "tit":
		return Titus, nil
	case "phlm":
		return Philemon, nil
	case "hebr":
		return Hebraer, nil
	case "jakbr":
		return Jakobus, nil
	case "1 petr":
		return ErstePetrus, nil
	case "2 petr":
		return ZweitePetrus, nil
	case "1 joh":
		return ErsteJohannes, nil
	case "2 joh":
		return ZweiteJohannes, nil
	case "3 joh":
		return DritteJohannes, nil
	case "jud":
		return Judas, nil
	case "offb":
		return Offenbarung, nil
	case "1 ne":
		return ErsteNephi, nil
	case "2 ne":
		return ZweiteNephi, nil
	case "jak":
		return Jakob, nil
	case "enos":
		return Enos, nil
	case "jar":
		return Jarom, nil
	case "om":
		return Omni, nil
	case "wmorm":
		return WorteMormons, nil
	case "mos":
		return Mosia, nil
	case "al":
		return Alma, nil
	case "hel":
		return Heleman, nil
	case "3 ne":
		return DritteNephi, nil
	case "4 ne":
		return VierteNephi, nil
	case "morm":
		return Mormon, nil
	case "eth":
		return Ether, nil
	case "moro":
		return Moroni, nil
	case "lub":
		return LehreUndBuendnisse, nil
	case "ae---1":
		return AmtlicheErklaerung1, nil
	case "ae---2":
		return AmtlicheErklaerung2, nil
	case "mose":
		return Mose, nil
	case "abr":
		return Abraham, nil
	case "jsmt":
		return JosephSmithMatthaeus, nil
	case "jslg":
		return JosephSmithLebensgeschichte, nil
	case "ga":
		return Glaubensartikel, nil
	default:
		return -1, errors.Errorf("unrecognized book: %s", s)
	}
}

func (b Book) String() string {
	switch b {
	case Genesis:
		return "Gen"
	case Exodus:
		return "Ex"
	case Levitikus:
		return "Lev"
	case Numeri:
		return "Num"
	case Deuteronomium:
		return "Dtn"
	case Josua:
		return "Jos"
	case Richter:
		return "Ri"
	case Rut:
		return "Rut"
	case ErsteSamuel:
		return "1 Sam"
	case ZweiteSamuel:
		return "2 Sam"
	case ErsteKoenige:
		return "1 Kön"
	case ZweiteKoenige:
		return "2 Kön"
	case ErsteChronik:
		return "1 Chr"
	case ZweiteChronik:
		return "2 Chr"
	case Esra:
		return "Esra"
	case Nehemia:
		return "Neh"
	case Ester:
		return "Est"
	case Ijob:
		return "Ijob"
	case Psalmen:
		return "Ps"
	case Sprichwoerter:
		return "Spr"
	case Kohelet:
		return "Koh"
	case Hohelied:
		return "Hld"
	case Jesaja:
		return "Jes"
	case Jeremia:
		return "Jer"
	case Klagelieder:
		return "Klgl"
	case Ezechiel:
		return "Ez"
	case Daniel:
		return "Dan"
	case Hosea:
		return "Hos"
	case Joel:
		return "Joël"
	case Amos:
		return "Am"
	case Obadja:
		return "Obd"
	case Jona:
		return "Jona"
	case Micha:
		return "Mi"
	case Nahum:
		return "Nah"
	case Habakuk:
		return "Hab"
	case Zefanja:
		return "Zef"
	case Haggai:
		return "Hag"
	case Sacharja:
		return "Sach"
	case Maleachi:
		return "Mal"
	case Mattaeus:
		return "Mt"
	case Markus:
		return "Mk"
	case Lukas:
		return "Lk"
	case Johannes:
		return "Joh"
	case Apostelgeschichte:
		return "Apg"
	case Roemer:
		return "Röm"
	case ErsteKorinther:
		return "1 Kor"
	case ZweiteKorinther:
		return "2 Kor"
	case Galater:
		return "Gal"
	case Epheser:
		return "Eph"
	case Philipper:
		return "Phil"
	case Kolosser:
		return "Kol"
	case ErsteThessalonicher:
		return "1 Thess"
	case ZweiteThessalonicher:
		return "2 Thess"
	case ErsteTimotheus:
		return "1 Tim"
	case ZweiteTimotheus:
		return "2 Tim"
	case Titus:
		return "Tit"
	case Philemon:
		return "Phlm"
	case Hebraer:
		return "Hebr"
	case Jakobus:
		return "Jakbr"
	case ErstePetrus:
		return "1 Petr"
	case ZweitePetrus:
		return "2 Petr"
	case ErsteJohannes:
		return "1 Joh"
	case ZweiteJohannes:
		return "2 Joh"
	case DritteJohannes:
		return "3 Joh"
	case Judas:
		return "Jud"
	case Offenbarung:
		return "Offb"
	case ErsteNephi:
		return "1 Ne"
	case ZweiteNephi:
		return "2 Ne"
	case Jakob:
		return "Jak"
	case Enos:
		return "Enos"
	case Jarom:
		return "Jar"
	case Omni:
		return "Om"
	case WorteMormons:
		return "WMorm"
	case Mosia:
		return "Mos"
	case Alma:
		return "Al"
	case Heleman:
		return "Hel"
	case DritteNephi:
		return "3 Ne"
	case VierteNephi:
		return "4 Ne"
	case Mormon:
		return "Morm"
	case Ether:
		return "Eth"
	case Moroni:
		return "Moro"
	case LehreUndBuendnisse:
		return "LuB"
	case AmtlicheErklaerung1:
		return "AE---1"
	case AmtlicheErklaerung2:
		return "AE---2"
	case Mose:
		return "Mose"
	case Abraham:
		return "Abr"
	case JosephSmithMatthaeus:
		return "JSMt"
	case JosephSmithLebensgeschichte:
		return "GSLg"
	case Glaubensartikel:
		return "GA"
	default:
		panic(errors.Errorf("illegal book value: %d", b))
	}
}

func (b Book) MarshalYAML() (interface{}, error) {
	return b.String(), nil
}

func (b *Book) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var raw string
	err := unmarshal(&raw)
	if err != nil {
		return errors.WithMessagef(err, "failed to unmarshal string while parsing book")
	}

	v, err := ParseBook(raw)
	if err != nil {
		return errors.WithStack(err)
	}

	*b = v
	return nil
}
