##程序员的数学
    此工程的代码，基于平冈和幸的《程序员的数学》系列一书的例子进行编码(原书的代码是由R语言编写的)。
    意图通过go语言把书籍中通俗易懂的数学思想表达出来。
###monty
    蒙提霍尔问题。三门问题：参赛者会看见三扇关闭了的门，其中一扇的后面有一辆汽车，选中后面有车的那扇门就可以赢得该汽车，
    而另外两扇门后面则各藏有一只山羊。当参赛者选定了一扇门，但未去开启它的时候，节目主持人会开启剩下两扇门的其中一扇，
    露出其中一只山羊。主持人其后会问参赛者要不要换另一扇仍然关上的门。问题是：换另一扇门会否增加参赛者赢得汽车的机会率？
    如果严格按照上述的条件的话，答案是会—换门的话，赢得汽车的机会率是 2/3。
    本例中采用是多协程模拟指定场数的形式，并将改变与不改变的结果，存入到最终数组，通过plot的方式，绘制到图上。
    其实三门问题，是多门问题的简化版本，意图是，当你不知道正确结果的时候，随机选择了一个结果，但是有外部力量
    把一个错误结果展示给你的时候，你应该作出改变还是不改变，这里不涉及任何有助于猜想其他情况的信息。正解是作出改变
    参见 https://zhuanlan.zhihu.com/p/464351914
    假设有门A、B、C
    门A后有将的概率事件a，概率为P(a)=1/3;
    门B后有将的概率事件b，概率为P(b)=1/3;
    门B后有将的概率事件c，概率为P(c)=1/3;
    到此是一个概率和为1的。
    参赛者随机选择一个门，假设为A门，是一个必然事件h
    主持人打开一个门，假设是B门，这是概率事件d P(d)
    求解的就是“空门打开后，不更改选择的胜率”P(a|d)，以及“更改选择的胜率”P(c|d)
    根据贝叶斯公式，P(a|d)=P(d|a)*P(a)/P(d),
    根据全概率公式，P(d)= P(d|a)*P(a)+ P(d|b)*P(b)+ P(d|c)*P(c)。
    P(c|d)=P(d|c)*P(c)/P(d)    

    条件概率P(d|a) 的物理意义是“当a事件，即A门后有大奖的事件发生时，主持人打开B门的概率“。 
    显然，这时主持人可以打开B门，也可以开C门，所以这个概率就是P(d|a)=1/2
    
    条件概率P(d|b) 的物理意义是“当b事件，即B门后有大奖的事件发生时，主持人打开B门的概率“。 
    这时主持人不能打开B门，所以这个概率就是P(d|b)=0。
    
    条件概率P(d|c) 的物理意义是“当c事件，即C门后有大奖的事件发生时，主持人打开B门的概率“。 
    这时主持人只能打开B门，所以这个概率就是P(d|c)=1。

    把上面已知数据代入，就得到“不更改的胜率“P(a|d)=1/2*1/3/(1/2*1/3+0*1/3+1*1/3)=1/3；
    而”更改后的胜率“P(c|d)=1*1/3/(1/2*1/3+0*1/3+1*1/3)=2/3。

    通过数学公式得到的结果会更有信服力。