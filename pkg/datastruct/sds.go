package datastruct

type SDS interface {
	SdsNew()
	SdsEmpty()
	SdsFree()
	SdsLen()
	SdsClear()
	SdsCat()
	SdsCpy()
	SdsRange()
	SdsTrim()
	SdsCmp()
}
