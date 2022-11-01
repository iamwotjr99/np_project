import 'react-native-gesture-handler';
import { NavigationContainer } from '@react-navigation/native';
import { createStackNavigator } from '@react-navigation/stack';
import Login from './component/Login';
import Board from './component/Board';
import Board_create from './component/Board_create';
import Menu from './component/Menu';
import Home from './component/Home'
import { verticalAnimation } from './animations/verticalAnimation';
import { NativeBaseProvider } from "native-base";


const Stack = createStackNavigator();

export default function App() {
  return (
    <NativeBaseProvider>
      <NavigationContainer>
        <Stack.Navigator>
          <Stack.Screen name="Login" component={Login} options={{headerShown: false}}/>
          {/* <Stack.Screen name="Board" component={Board} options={{title: '게시판'}}/> */}
          <Stack.Screen name="Board_create" component={Board_create} options={verticalAnimation} />
          {/* <Stack.Screen name="Home" component={Home} options={{headerShown: false}} /> */}
          <Stack.Screen name='BottomMenu' component={Menu} options={{title: "Study with me"}}/>
        </Stack.Navigator>
      </NavigationContainer>
    </NativeBaseProvider>
  );
}

