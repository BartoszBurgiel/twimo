// In React Native, the component exported from App.js is the App's entry point (root component)
// So, here we have to define the routing and screens, like which screen is when & where displayed

// Core functionality imports for navigation
import { createAppContainer } from "react-navigation";
import { createStackNavigator } from "react-navigation-stack";

// Component & screen imports
import HomeScreen from "./src/screens/HomeScreen";
import DetailsScreen from "./src/screens/DetailsScreen";
import ListScreen from "./src/screens/ListScreen";

// Define the AppNavigator structure
const AppNavigator = createStackNavigator(
  {
    Home: HomeScreen,
    Details: DetailsScreen,
    List: ListScreen
  },
  {
    initialRouteName: "Home"
  }
);

// Export our just defined root component
export default createAppContainer(AppNavigator);
