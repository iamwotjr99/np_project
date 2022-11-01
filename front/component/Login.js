import {View, StyleSheet, Image} from 'react-native'
import {useState} from 'react'
import { Box, FormControl, Stack, Input, Button } from 'native-base'
import axios from 'axios'

const BASE_URL = "http://10.10.1.168:80/api"

// http://192.168.25.52:80
// http://192.168.0.23:80/api

function Login({navigation}) {
    const [name, setName] = useState(null)
    const [pw, setPw] = useState(null)
    
    // const getPressBtn = async () => {
    //     await axios.get(`${BASE_URL}/api/someGet`).then((res) => {
    //         console.log(res.data);
    //     })
    // }

    const loginPressBtn = async () => {
        await axios.post(`${BASE_URL}/user/auth`, {
            name: name,
            pw: pw
        }).then((res) => {
            if(res.status == 200) {
                navigation.navigate('BottomMenu', {name: name})
            }
            setName(null)
            setPw(null)
        })
    }

    return (
        <View style={styles.container}>
            <Box alignItems="center">
                <Image source={require('../assets/logo.png')} style={{width: 400, resizeMode: 'contain'}}/>
                <Box w="100%" maxHeight="300px">
                    <FormControl marginBottom={"2%"}>
                        <Stack mx="6">
                            <FormControl.Label fontSize={"lg"}>이름</FormControl.Label>
                            <Input size="lg" type="text" placeholder="name" onChangeText={setName} value={name}/>
                        </Stack>
                    </FormControl>
                    <FormControl marginBottom={"2%"}>
                        <Stack mx="6">
                            <FormControl.Label>비밀번호</FormControl.Label>
                            <Input size="lg" type="password" placeholder="password" onChangeText={setPw} value={pw}/>
                        </Stack>
                    </FormControl>
                </Box>
                <Button onPress={loginPressBtn} style={styles.button}>Login</Button>
            </Box>
        </View>
    )
}

export default Login;

const styles = StyleSheet.create({
    container: {
        flex: 1,
        alignContent: "center",
        justifyContent: 'center',
        padding: "5%",
        backgroundColor: "white"
    },

    button: {
        marginTop: "5%",
        width: "50%"
    }
});