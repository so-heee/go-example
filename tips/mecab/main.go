package main

import (
	"fmt"
	"strings"

	"github.com/bluele/mecab-golang"
)

const BOSEOS = "BOS/EOS"

func parseToNode(m *mecab.MeCab, text string) error {

	tg, err := m.NewTagger()
	if err != nil {
		return err
	}
	defer tg.Destroy()

	lt, err := m.NewLattice(text)
	if err != nil {
		return err
	}

	node := tg.ParseToNode(lt)
	for {
		features := strings.Split(node.Feature(), ",")
		if features[0] != BOSEOS {
			fmt.Printf("%s %s\n", node.Surface(), node.Feature())
		}
		if node.Next() != nil {
			break
		}
	}
	return nil
}

func main() {
	m, err := mecab.New("-Owakati")
	if err != nil {
		panic(err)
	}
	defer m.Destroy()

	err = parseToNode(m, "吾輩は猫である。名前はまだ無い。どこで生れたかとんと見当けんとうがつかぬ。何でも薄暗いじめじめした所でニャーニャー泣いていた事だけは記憶している。吾輩はここで始めて人間というものを見た。しかもあとで聞くとそれは書生という人間中で一番獰悪どうあくな種族であったそうだ。この書生というのは時々我々を捕つかまえて煮にて食うという話である。しかしその当時は何という考もなかったから別段恐しいとも思わなかった。ただ彼の掌てのひらに載せられてスーと持ち上げられた時何だかフワフワした感じがあったばかりである。掌の上で少し落ちついて書生の顔を見たのがいわゆる人間というものの見始みはじめであろう。この時妙なものだと思った感じが今でも残っている。第一毛をもって装飾されべきはずの顔がつるつるしてまるで薬缶やかんだ。その後ご猫にもだいぶ逢あったがこんな片輪かたわには一度も出会でくわした事がない。のみならず顔の真中があまりに突起している。そうしてその穴の中から時々ぷうぷうと煙けむりを吹く。どうも咽むせぽくて実に弱った。これが人間の飲む煙草たばこというものである事はようやくこの頃知った。")
	if err != nil {
		panic(err)
	}
}
