import { createBottomTabNavigator } from '@react-navigation/bottom-tabs'
import { NavigationContainer } from '@react-navigation/native';
import Icon from 'react-native-vector-icons/FontAwesome'
import Board from './Board'
import Home from './Home'
import Chat from './Chat'
import User from './User'
function Menu({navigation, route}) {
    const Tab = createBottomTabNavigator()

    const {name} = route.params

    return(
        
            <Tab.Navigator initialRouteName='Board' screenOptions={({ route }) => ({
                headerShown: false,
                tabBarIcon: ({focused, color, size}) => {
                    let iconName
                    if(route.name === 'Home') {
                        iconName = "home"
                        color= focused ? "rgb(90, 173, 178)" : "black"
                    }
                    else if(route.name === 'Board') {
                        iconName = "clipboard"
                        color= focused ? "rgb(90, 173, 178)" : "black"
                    }
                    else if(route.name === 'Chat') {
                        iconName = "wechat"
                        color= focused ? "rgb(90, 173, 178)" : "black"
                    }
                    else if(route.name === 'User') {
                        iconName = "user"
                        color= focused ? "rgb(90, 173, 178)" : "black"
                    }
                    size=30
                    return <Icon name={iconName} size={size} color={color} />
                }
            })}>
                <Tab.Screen name="Home" children={({navigation})=><Home navigation={navigation} name={name}/>} options={{tabBarShowLabel: false}}/>
                <Tab.Screen name="Board" children={({navigation})=><Board navigation={navigation} name={name}/>} options={{tabBarShowLabel: false}}/>
                <Tab.Screen name="Chat" children={({navigation})=><Chat navigation={navigation} name={name}/>} options={{tabBarShowLabel: false}}/>
                <Tab.Screen name="User" children={({navigation})=><User navigation={navigation} name={name}/>} options={{tabBarShowLabel: false}}/>
            </Tab.Navigator>
        
    )
}

export default Menu