package main

import (

	"strconv"
	"fyne.io/fyne/v2"
	// "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Knetic/govaluate"
)

func showCalc() {
	// a := app.New()
	// w := a.NewWindow("Calculator")

	output:=""
	var ArrayOfHistory[] string;
	 flag:=false;
	input := widget.NewLabel(output)

	historyStr:="";
	history:=widget.NewLabel(historyStr)

	historyBtn:= widget.NewButton("History",func(){

		if flag==false{

			for i:= len(ArrayOfHistory)-1; i>=0; i--{
				historyStr = historyStr+ArrayOfHistory[i];
				historyStr +="\n";
			}
			history.SetText(historyStr);
		}else{
			historyStr="";
		}

		flag = !flag;
	})
	backBtn:= widget.NewButton("Back",func(){

		if len(output)>0{
			output =  output[:len(output)-1];
			input.SetText(output)}
		
	})
	clearBtn:= widget.NewButton("Clear",func(){
		
		output="";
		historyStr="";
		input.SetText(output);
		history.SetText(historyStr);
	})
	lbrackerBtn:= widget.NewButton("(",func(){
		
		output = output+"(";
		input.SetText(output);
	})
	rbrackerBtn:= widget.NewButton(")",func(){

		output = output+")";
		input.SetText(output);
	
	})
	slashBtn:= widget.NewButton("/",func(){
		
		output = output+"/";
		input.SetText(output);
	})
	sevenBtn:= widget.NewButton("7",func(){
		
		output = output+"7";
		input.SetText(output);
	})
	eightBtn:= widget.NewButton("8",func(){
		
		output = output+"8";
		input.SetText(output);
	})
	nineBtn:= widget.NewButton("9",func(){
		
		output = output+"9";
		input.SetText(output);
	})
	starBtn:= widget.NewButton("*",func(){
		
		output = output+"*";
		input.SetText(output);
	})
	fourBtn:= widget.NewButton("4",func(){
		
		output = output+"4";
		input.SetText(output);
	})
	fiveBtn:= widget.NewButton("5",func(){
		
		output = output+"5";
		input.SetText(output);
	})
	sixBtn:= widget.NewButton("6",func(){
		
		output = output+"6";
		input.SetText(output);
	})
	minusBtn:= widget.NewButton("-",func(){
		
		output = output+"-";
		input.SetText(output);
	})
	oneBtn:= widget.NewButton("1",func(){
		
		output = output+"1";
		input.SetText(output);
	})
	twoBtn:= widget.NewButton("2",func(){
		
		output = output+"2";
		input.SetText(output);
	})
	threeBtn:= widget.NewButton("3",func(){
		
		output = output+"3";
		input.SetText(output);
	})
	plusBtn:= widget.NewButton("+",func(){
		
		output = output+"+";
		input.SetText(output);
	})
	zeroBtn:= widget.NewButton("0",func(){
		
		output = output+"0";
		input.SetText(output);
	})
	dotBtn:= widget.NewButton(".",func(){
		output = output+".";
		input.SetText(output);
	})
	equalBtn:= widget.NewButton("=",func(){
		
		expression, err := govaluate.NewEvaluableExpression(output);
		if err==nil{
			result, err := expression.Evaluate(nil);

			if err==nil{
				ans:= strconv.FormatFloat(result.(float64),'f',-1,64);
				strToAppend:= output+"="+ans;
				ArrayOfHistory = append(ArrayOfHistory, strToAppend);
				output=ans;
			}else{
				output="error";
			}
		}else{
			output="error";
		}

		input.SetText(output);
		
	})


	calcContainer:=container.NewVBox(
		input,
		history,
		container.NewGridWithColumns(1,
			
			container.NewGridWithColumns(2,
				historyBtn,
				backBtn,
			),

			container.NewGridWithColumns(4,
				clearBtn,
				lbrackerBtn,
				rbrackerBtn,
				slashBtn,
			),
			container.NewGridWithColumns(4,
				sevenBtn,
				eightBtn,
				nineBtn,
				starBtn,
			),
			container.NewGridWithColumns(4,
				fourBtn,
				fiveBtn,
				sixBtn,
				minusBtn,
			),
			container.NewGridWithColumns(4,
				oneBtn,
				twoBtn,
				threeBtn,
				plusBtn,
			),
			container.NewGridWithColumns(2,
				zeroBtn,
				dotBtn,
			),
			container.NewGridWithColumns(1,
				equalBtn,
			),
		),
	)


	w:=myApp.NewWindow("calculator")	
	w.Resize(fyne.NewSize(500,280));

	w.SetContent(
		container.NewBorder(DeskBtn,nil,nil,nil,calcContainer),
	)	
	w.Show();	
	// w.ShowAndRun()
}