// In React Native, the component exported from App.js is the App's entry point

// Core functionality imports for navigation
import { createAppContainer } from "react-navigation";
import { createStackNavigator } from "react-navigation-stack";

// Component & screen imports
import HomeScreen from "./src/screens/HomeScreen";

const AppNavigator = createStackNavigator({
  Home: {
    screen: HomeScreen
  }
});

export default createAppContainer(AppNavigator);
