import { View, Text, StyleSheet } from 'react-native'
import { useState, useEffect } from 'react'
import axios from 'axios'
import { useIsFocused } from '@react-navigation/native'
import Board_item from './Board_item'

const BASE_URL = "http://10.10.1.168:80/api"

function Board({navigation, name}) {
    const isFocused = useIsFocused()

    const [articleList, setList] = useState([]);

    async function getArticleList() {
        await axios.get(`${BASE_URL}/articles`).then((res) => {
            let result = res.data.result
            setList(result)
        })
    }

    // async function makeSyncList() {
    //     articleList = await getArticleList()
    //     // console.log("articleList:", articleList)
    // }

    useEffect(() => {
        getArticleList()
    }, [isFocused])

    const getDiffMin = (createdAt) => {
        const _now = new Date();
        _now.setHours(_now.getHours() + 9)
        const now = _now.toISOString().replace('T', ' ').substring(0, 19)
        const past = createdAt.replace(/-/g, '/')
        // console.log(past)
        let r_past = new Date(createdAt)
        let r_now = new Date(now)
        let diffMin = (r_now.getTime() - r_past.getTime()) / (1000*60);
        // console.log(r_past)
        return diffMin
    }

    function textLengthOverCut(txt, len) {
        if(txt.length > len) {
            txt = txt.substr(0, len) + '...'
        }

        return txt
    }

    const onPressBtn = () => {
        navigation.navigate('Board_create', {name: name})
    }

    return (
        <View style={styles.container}>
            {articleList && articleList.map((item, key) => {
                const diffMin = getDiffMin(item.createdAt)
                return (<Board_item 
                            id={item.id} 
                            author={item.author} 
                            title={item.title}
                            content={textLengthOverCut(item.content, 34)}
                            diffMin={diffMin}
                            key={key}
                        />)
            })}
            <View style={styles.floating_btn_wrap}>
                <Text style={styles.floating_btn} onPress={onPressBtn}>✏️  글쓰기</Text>
            </View>
        </View>
    )
}

export default Board

const styles = StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: "white",
        padding: "4%",
    },

    floating_btn_wrap: {
        alignItems: 'center',
        marginTop: 'auto',                                              
    },

    floating_btn: {
        borderWidth: 1,
        padding: '4%',
        borderRadius: 20,
        borderColor: "rgba(0, 0, 0, 0.3)"
    }
});