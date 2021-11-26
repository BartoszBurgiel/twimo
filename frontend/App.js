// In React Native, the component exported from App.js is the App's entry point (root component)
// So, here we have to define the routing and screens, like which screen is when & where displayed

// Core functionality imports for navigation
import { createAppContainer } from "react-navigation";
import { createStackNavigator } from "react-navigation-stack";

// Screen imports
import HomeScreen from "./src/screens/HomeScreen";
import ListScreen from "./src/screens/ListScreen";
import LocationDetailsScreen from "./src/screens/LocationDetailsScreen";
import SignUpScreen from "./src/screens/SignUpScreen";

// Define the AppNavigator structure
const AppNavigator = createStackNavigator(
  {
    Home: HomeScreen,
    List: ListScreen,
    LocationDetails: LocationDetailsScreen,
    SignUp: SignUpScreen
  },
  {
    initialRouteName: "Home"
  }
);

// Export our just defined root component
export default createAppContainer(AppNavigator);
