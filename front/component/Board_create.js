import { View, Text, StyleSheet, StatusBar, Platform } from 'react-native'
import { getStatusBarHeight } from 'react-native-status-bar-height';
import { TextInput } from 'react-native-gesture-handler';
import { useState } from 'react';
import axios from 'axios';

const StatusBarHeight =
    Platform.OS === 'ios' ? getStatusBarHeight(true) : StatusBar.currentHeight;

const BASE_URL = "http://10.10.1.168:80/api"

function Board_create({navigation : { goBack }, route}) {
    const {name} = route.params
    const [article, setArticle] = useState({
        author: name,
        title: "",
        content: "",
        height: 0
    })

    const changeHandler = (name, e) => {
        setArticle({
            ...article,
            [name]: e,
        })
    }

    async function postArticle() {
        await axios.post(`${BASE_URL}/articles`, {
            author: article.author,
            title: article.title,
            content: article.content
        }).then((res) => {
            console.log(res.data)
            goBack()
        })
    } 

    return (
        <View style={styles.container}>
            <View style={styles.header}>
                <Text style={styles.cancle} 
                    onPress={() => goBack() }>X</Text>
                <Text style={styles.header_title}>글 쓰기</Text>
                <Text style={styles.right} onPress={postArticle}>완료</Text>
            </View>
            <View style={{flex: 1}}>
                <View>
                    <TextInput style={styles.title} 
                        placeholder="제목"
                        onChangeText={(e) => changeHandler('title', e)}
                        value={article.title}/>
                </View>
                <TextInput style={[styles.textarea, {height: Math.max(35, article.height)}]}
                    multiline={true}
                    placeholder="내용을 입력하세요."
                    onChangeText={(e) => changeHandler('content', e)}
                    onContentSizeChange={(e) => {setArticle({...article, height: e.nativeEvent.contentSize.height})}}
                    value={article.content}
                    />
            </View>
        </View>
    )
}

export default Board_create

const styles = StyleSheet.create({
    container: {
        padding: "6%",
        backgroundColor: "white",
        flex: 1,
    },
    header: {
        flexDirection: 'row',
        marginTop: StatusBarHeight,
        paddingBottom: '2%',
        marginBottom: "5%"
    },
    cancle: {
        fontSize: 20,
        marginRight: "8%"
    },
    header_title: {
        fontSize: 20,
        fontWeight: 'bold'
    },
    right: {
        fontSize: 20,
        marginLeft: 'auto'
    },

    title: {
        fontSize: 19,
        fontWeight: 'bold',
        height: 50,
        borderBottomWidth: 1,
        borderBottomColor: "rgba(0, 0, 0, 0.3)",
        marginBottom: "2%",
    },
    textarea: {
        justifyContent: "flex-start",
        textAlignVertical: 'top',
        flex: 1,
        fontSize: 15
    }
    
})