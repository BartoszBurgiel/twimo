// In React Native, the component exported from App.js is the App's entry point

// Core functionality imports for navigation
import { createAppContainer } from "react-navigation";
import { createStackNavigator } from "react-navigation-stack";

// Component & screen imports
import HomeScreen from "./src/screens/HomeScreen";
import DetailsScreen from "./src/screens/DetailsScreen";
import ListScreen from "./src/screens/ListScreen";

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

export default createAppContainer(AppNavigator);
